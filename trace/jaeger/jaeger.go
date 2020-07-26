package jaeger

/*

ref:
	- https://github.com/jaegertracing/jaeger-client-go

*/

import (
	"io"
	"time"

	"github.com/better-go/pkg/log"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

// New 创建一个 jaeger Tracer
func New(serviceName string, addr string) (opentracing.Tracer, io.Closer, error) {
	cfg := config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	// conn Jaeger Server:
	sender, err := jaeger.NewUDPTransport(addr, 0)
	if err != nil {
		log.Errorf("new jaeger sender error, err=%v", err)
		return nil, nil, err
	}

	reporter := jaeger.NewRemoteReporter(sender)

	// Initialize tracer with a logger and a metrics factory
	tracer, closer, err := cfg.NewTracer(
		config.Reporter(reporter),
	)
	return tracer, closer, err
}
