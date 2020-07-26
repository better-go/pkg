package zipkin

import (
	"github.com/better-go/pkg/log"
	"github.com/opentracing/opentracing-go"
	zipkinTrace "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zipkinHttp "github.com/openzipkin/zipkin-go/reporter/http"
)

/*

ref:
	- https://github.com/openzipkin-contrib/zipkin-go-opentracing
	- https://github.com/openzipkin/zipkin
	- https://github.com/openzipkin/zipkin-go

*/

// New 创建一个 jaeger Tracer
func New(serviceName string, serviceAddr string, addr string) (opentracing.Tracer, error) {
	reporterUrl := "http://zipkinhost:9411/api/v2/spans"
	if addr != "" {
		reporterUrl = addr + "/api/v2/spans"
	}

	// set up a span reporter
	reporter := zipkinHttp.NewReporter(reporterUrl)
	defer reporter.Close()

	//create our local service endpoint
	endpoint, err := zipkin.NewEndpoint(serviceName, "")
	if err != nil {
		log.Errorf("trace.zipkin: create local endpoint error: %+v", err)
		return nil, err
	}

	// new tracer
	nakedT, err := zipkin.NewTracer(
		reporter,
		zipkin.WithLocalEndpoint(endpoint),
		zipkin.WithNoopTracer(true),
	)
	log.Infof("trace.zipkin: create tracer: trace=%v, err=%v", nakedT, err)
	if err != nil {
		return nil, err
	}

	// use zipkin-go-opentracing to wrap our t
	t := zipkinTrace.Wrap(nakedT)

	// optionally set as Global OpenTracing t instance
	opentracing.SetGlobalTracer(t)
	return t, nil
}
