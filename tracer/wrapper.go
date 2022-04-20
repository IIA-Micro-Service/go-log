package tracer

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	resource2 "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"log"
)

var globalTracerProvider *sdktrace.TracerProvider
var globalTracer trace.Tracer

/*
 * @desc : 初始化一个exporter
 */
func newExporter() *stdout.Exporter {
	exporter, err := stdout.New(stdout.WithPrettyPrint())
	if err != nil {
		log.Fatal("Init Exporter Err:%+v", err)
	}
	return exporter
}

/*
 * @desc : 初始化一个tracer provider
 */
func newTracerProvider(exporter sdktrace.SpanExporter) *sdktrace.TracerProvider {
	resource := resource2.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("example"),
	)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource),
	)
	otel.SetTracerProvider(tracerProvider)
	fmt.Printf("1:%+v\n", tracerProvider)
	return tracerProvider
}

func NewTracerWrapper() (*sdktrace.TracerProvider, trace.Tracer) {

	exporter := newExporter()
	globalTracerProvider = newTracerProvider(exporter)
	fmt.Printf("2:%+v\n", globalTracerProvider)
	//stdout.
	//tracer = tp.Tracer("ExampleService")
	globalTracer = globalTracerProvider.Tracer("example")

	_, span := globalTracer.Start(context.Background(), "example-span")
	//fmt.Printf("%+v\n", ctx)
	defer span.End()

	//fmt.Printf("%+v", span.SpanContext())
	return globalTracerProvider, globalTracer

}
