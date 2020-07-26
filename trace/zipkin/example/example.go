package main

import (
	"context"
	"errors"
	"time"

	pkgTrace "github.com/better-go/pkg/trace/zipkin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

/*

ref:
	- https://juejin.im/post/5d7ed711e51d4562165535ab

*/

func main() {
	callV1() // v1
	//callV2() // v2
}

func callV1() {
	// new trace:
	t, err := pkgTrace.New("com.app1.demo1", ":9411", "127.0.0.1:9411")
	if err != nil {
		return
	}
	opentracing.SetGlobalTracer(t)

	// 获取jaeger tracer
	//t := opentracing.GlobalTracer()

	// 创建root span
	sp := t.StartSpan("in-process-service")
	// main执行完结束这个span
	defer sp.Finish()
	// 将span传递给Foo
	ctx := opentracing.ContextWithSpan(context.Background(), sp)
	Foo(ctx)
}

func Foo(ctx context.Context) {
	// 开始一个span, 设置span的operation_name=Foo
	span, ctx := opentracing.StartSpanFromContext(ctx, "Foo")
	defer span.Finish()
	// 将context传递给Bar
	Bar(ctx)
	// 模拟执行耗时
	time.Sleep(1 * time.Second)
}
func Bar(ctx context.Context) {
	// 开始一个span，设置span的operation_name=Bar
	span, ctx := opentracing.StartSpanFromContext(ctx, "Bar")
	defer span.Finish()
	// 模拟执行耗时
	time.Sleep(2 * time.Second)

	// 假设Bar发生了某些错误
	err := errors.New("something wrong")
	span.LogFields(
		log.String("event", "error"),
		log.String("message", err.Error()),
	)
	span.SetTag("error", true)
}
