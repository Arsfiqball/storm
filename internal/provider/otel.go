package provider

import (
	"os"

	"go.opentelemetry.io/otel"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

type Otel interface {
	Provider() *sdktrace.TracerProvider
	Tracer() trace.Tracer
}

type otelState struct {
	provider *sdktrace.TracerProvider
	tracer   trace.Tracer
}

func ProvideOtel() (Otel, error) {
	var (
		sampler  sdktrace.Sampler
		exporter sdktrace.SpanExporter
		res      *resource.Resource
		err      error
	)

	if os.Getenv("ZIPKIN_URL") != "" {
		sampler = sdktrace.AlwaysSample()
		exporter, err = zipkin.New(os.Getenv("ZIPKIN_URL"))

		if err != nil {
			return nil, err
		}

		res = resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("my_app"),
		)
	} else {
		sampler = sdktrace.AlwaysSample()
		exporter, err = stdout.New()

		if err != nil {
			return nil, err
		}

		res = resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("my_app"),
		)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sampler),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	otel.SetTracerProvider(tp)

	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	return &otelState{
		provider: tp,
		tracer:   tp.Tracer("my_app"),
	}, nil
}

func (o *otelState) Provider() *sdktrace.TracerProvider {
	return o.provider
}

func (o *otelState) Tracer() trace.Tracer {
	return o.tracer
}
