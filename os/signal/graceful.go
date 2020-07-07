package signal

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/better-go/pkg/log"
)

// task start handler
type StartFunc func() error

// task clean handler
type StopFunc func() error

// graceful shutdown:
func GracefulShutdown(ctx context.Context, startFn StartFunc, stopFn StopFunc) error {
	//
	// start:
	//
	if err := startFn(); err != nil {
		log.Error("start error, err=%v", err)
		return err
	}

	log.Infof("graceful start...")

	// os signal chan:
	ch := make(chan os.Signal, 1)

	// register os exit signal:
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	//
	// watch os exit signal:
	//

	select {
	// wait on kill signal
	case <-ch:
	// wait on context cancel
	case <-ctx.Done():
	}

	//
	// now ready to exit, do clean
	//
	if err := stopFn(); err != nil {
		log.Error("shutdown error, err=%v", err)
		return err
	}

	log.Infof("graceful shutdown...")
	return nil
}
