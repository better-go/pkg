package async

import (
	"context"
	"errors"
	"runtime"
	"sync"

	pkgCtx "github.com/better-go/pkg/context"
	"github.com/better-go/pkg/log"
)

var (
	// ErrFull chan full.
	ErrFull = errors.New("async: chan full")
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////
//                              异步任务分发器:
//
// 说明:
//	- 基于 go channel 实现.
// 	- 操作异步化
//	- close() 方法, 必须显示在 main 进程调用. 防止进程强退, 丢 task
//
//////////////////////////////////////////////////////////////////////////////////////////////////////////

// TaskDispatcher async consume data from chan.
type TaskDispatcher struct {
	name    string // 分组标签
	ch      chan item
	options *options
	waiter  sync.WaitGroup

	ctx    context.Context
	cancel func()
}

// New new a TaskDispatcher struct.
func New(name string, opts ...Option) *TaskDispatcher {
	// label 分组:
	if name == "" {
		name = "anonymous"
	}

	// 默认配置:
	o := &options{
		worker: 1,
		buffer: 1024,
	}

	// 初始化:
	for _, op := range opts {
		op(o)
	}

	// 创建:
	c := &TaskDispatcher{
		ch:      make(chan item, o.buffer),
		name:    name,
		options: o,
	}

	// ctx:
	c.ctx, c.cancel = context.WithCancel(context.Background())
	c.waiter.Add(o.worker)

	//
	// 异步处理: 监听+接收+处理
	//
	for i := 0; i < o.worker; i++ {
		// 异步例程: 监听
		go c.taskHandleProc()
	}
	return c
}

// taskHandleProc taskFn 处理
func (c *TaskDispatcher) taskHandleProc() {
	defer c.waiter.Done()

	//
	// loop:
	//
	for {
		select {
		case task := <-c.ch: // receiver: 监听+接收+处理
			// do task:
			wrapHandleFunc(task.taskFn)(task.ctx)

		case <-c.ctx.Done(): // done:
			return
		}
	}
}

// wrapHandleFunc 运行时 check
func wrapHandleFunc(taskFn func(c context.Context)) (res func(context.Context)) {
	res = func(ctx context.Context) {
		defer func() {
			if r := recover(); r != nil {
				buf := make([]byte, 64*1024)

				buf = buf[:runtime.Stack(buf, false)]
				log.Errorf("panic in async.TaskDispatcher taskHandleProc, err: %s, stack: %s", r, buf)
			}
		}()

		//
		// call task func:
		//
		taskFn(ctx)
	}
	return
}

// Dispatch 分发调度 save a callback func.
func (c *TaskDispatcher) Dispatch(ctx context.Context, taskFn func(ctx context.Context)) (err error) {
	if taskFn == nil || c.ctx.Err() != nil {
		return c.ctx.Err()
	}

	// parse ctx:
	nakedCtx := pkgCtx.WithContext(ctx)

	//
	// 分发异步任务:
	//
	select {
	case c.ch <- item{taskFn: taskFn, ctx: nakedCtx}: // 分发
	default:
		err = ErrFull
	}
	return
}

// Close 必须显示调用, 防止丢 task
func (c *TaskDispatcher) Close() error {
	if err := c.ctx.Err(); err != nil {
		return err
	}
	c.cancel()

	//
	// 必须显示调用, 等待处理完成. 防止丢 task
	//
	c.waiter.Wait()
	return nil
}
