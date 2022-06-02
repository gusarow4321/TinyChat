package tracing

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"time"
)

func SetTracerProvider(url, name, id, version string) (func(log.Logger), error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(name),
			attribute.String("host", id),
			attribute.String("version", version),
		)),
	)
	otel.SetTracerProvider(tp)

	cleanup := func(logger log.Logger) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		if err := exp.Shutdown(ctx); err != nil {
			log.NewHelper(logger).WithContext(ctx).Errorf("Error while shutdown jaeger: %v", err)
		}
		cancel()
	}

	return cleanup, nil
}
