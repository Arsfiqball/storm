package provider

import (
	"strings"

	"github.com/spf13/viper"
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

	// Get configuration from viper
	serviceName := viper.GetString("telemetry.service_name")
	if serviceName == "" {
		serviceName = "storm-service"
	}

	samplingStrategy := viper.GetString("telemetry.sampling")
	zipkinURL := viper.GetString("telemetry.zipkin_url")

	// Configure sampler based on configuration
	switch strings.ToLower(samplingStrategy) {
	case "never":
		sampler = sdktrace.NeverSample()
	case "traceidratio":
		ratio := viper.GetFloat64("telemetry.sampling_ratio")
		if ratio <= 0 {
			ratio = 0.1 // default 10% sampling
		}
		sampler = sdktrace.TraceIDRatioBased(ratio)
	case "parentbased":
		sampler = sdktrace.ParentBased(sdktrace.AlwaysSample())
	default:
		// Default to AlwaysSample
		sampler = sdktrace.AlwaysSample()
	}

	if zipkinURL != "" {
		exporter, err = zipkin.New(zipkinURL)
		if err != nil {
			return nil, err
		}
	} else {
		exporter, err = stdout.New()

		if err != nil {
			return nil, err
		}
	}

	res = resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(serviceName),
	)

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
		tracer:   tp.Tracer(serviceName),
	}, nil
}

func (o *otelState) Provider() *sdktrace.TracerProvider {
	return o.provider
}

func (o *otelState) Tracer() trace.Tracer {
	return o.tracer
}
