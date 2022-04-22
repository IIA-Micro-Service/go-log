package main

import (
	"context"
	"github.com/IIA-Micro-Service/go-log"
	"github.com/IIA-Micro-Service/go-log/adapter"
	"github.com/IIA-Micro-Service/go-log/config"
	"github.com/IIA-Micro-Service/go-log/tracer"
)

func main() {

	logConfig := &config.Config{
		Type:         "logrus",
		LogDir:       "./",
		LogFileName:  "logrus.log",
		MaxLife:      7 * 86400,
		RotationTime: 1 * 86400,
		LogLevel:     adapter.TraceLevel,
	}
	loggerHandler := log.Init(logConfig)

	// 利用tracer生成span...假装是一次正常请求
	globalTracer := tracer.GetGlobalTracer()
	ctx, span := globalTracer.Start(context.Background(), "example-span")
	//fmt.Printf("%+v\n", ctx)
	defer span.End()
	//spanCtx := span.SpanContext()

	loggerHandler.WithContext(ctx).WithFields(adapter.Fields{
		"traceid": "",
		"spanid":  "",
		//"trace": time.Now(),
	}).Info("xi")

	//tp := otel.GetTracerProvider()
	//fmt.Printf("3:%+v\n", tp)

}
