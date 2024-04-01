package otel

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"go-boilerplate/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type Tracer struct {
	exporter *otlptrace.Exporter
	resource *resource.Resource
}

func NewTracer(config *config.Config) *Tracer {
	exporter, err := otlptrace.New(
		context.Background(),
		otlptracehttp.NewClient(
			otlptracehttp.WithEndpoint(config.Otel.CollectorUrl),
			otlptracehttp.WithInsecure(),
		),
	)

	if err != nil {
		log.Error().Err(err).Msg("Failed to create resource")
	}

	resource := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(config.Otel.ServiceName),
	)

	if err != nil {
		log.Error().Err(err).Msg("Failed to create resource")
	}

	return &Tracer{
		exporter: exporter,
		resource: resource,
	}
}

func (s *Tracer) Init() {
	otel.SetTracerProvider(
		sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithBatcher(s.exporter),
			sdktrace.WithResource(s.resource),
		),
	)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
}

func Start(ctx context.Context, name string, otparams map[string]interface{}) (context.Context, oteltrace.Span) {
	kv := []attribute.KeyValue{}
	for key, value := range otparams {
		jsonValue, _ := json.Marshal(value)
		kv = append(kv, attribute.String(key, fmt.Sprintf("%s", jsonValue)))
	}

	tracer := otel.Tracer("tracer")

	return tracer.Start(ctx, name, oteltrace.WithAttributes(kv...))
}
