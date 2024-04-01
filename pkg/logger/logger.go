package logger

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

func Info(ctx context.Context) *zerolog.Event {
	span := trace.SpanFromContext(ctx)
	return log.Info().Str("trace_id", span.SpanContext().TraceID().String()).Str("span_id", span.SpanContext().SpanID().String())
}

func Error(ctx context.Context) *zerolog.Event {
	span := trace.SpanFromContext(ctx)
	return log.Error().Str("trace_id", span.SpanContext().TraceID().String()).Str("span_id", span.SpanContext().SpanID().String())
}

func Warn(ctx context.Context) *zerolog.Event {
	span := trace.SpanFromContext(ctx)
	return log.Warn().Str("trace_id", span.SpanContext().TraceID().String()).Str("span_id", span.SpanContext().SpanID().String())
}

func Debug(ctx context.Context) *zerolog.Event {
	span := trace.SpanFromContext(ctx)
	return log.Debug().Str("trace_id", span.SpanContext().TraceID().String()).Str("span_id", span.SpanContext().SpanID().String())
}
