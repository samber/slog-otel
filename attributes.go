package slogcommon

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"
)

func ExtractOtelAttrFromContext(groups []string, traceIDKey string, spanIDKey string) func(ctx context.Context) []slog.Attr {
	return func(ctx context.Context) []slog.Attr {
		span := trace.SpanFromContext(ctx)
		if !span.IsRecording() {
			return []slog.Attr{}
		}

		attrs := []slog.Attr{}
		spanCtx := span.SpanContext()

		if spanCtx.HasTraceID() {
			traceID := trace.SpanFromContext(ctx).SpanContext().TraceID().String()
			attrs = append(attrs, slog.String(traceIDKey, traceID))
		}

		if spanCtx.HasSpanID() {
			spanID := spanCtx.SpanID().String()
			attrs = append(attrs, slog.String(spanIDKey, spanID))
		}

		for len(groups) > 0 {
			attrs = []slog.Attr{slog.Any(groups[len(groups)-1], attrs)}
			groups = groups[:len(groups)-1]
		}

		return attrs
	}
}
