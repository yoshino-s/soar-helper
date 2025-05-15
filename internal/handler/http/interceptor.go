package http

import (
	"context"

	"connectrpc.com/connect"
	"github.com/yoshino-s/go-framework/handlers/http"
	"go.opentelemetry.io/otel/semconv/v1.20.0/httpconv"
	"go.opentelemetry.io/otel/trace"
)

var _ connect.Interceptor = (*interceptor)(nil)

type interceptor struct{}

// WrapStreamingClient implements connect.Interceptor.
func (i *interceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return next
}

// WrapStreamingHandler implements connect.Interceptor.
func (i *interceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return func(ctx context.Context, shc connect.StreamingHandlerConn) error {
		echoCtx, _ := http.EchoContextFromContext(ctx)
		span := trace.SpanFromContext(ctx)
		span.SetAttributes(
			httpconv.ServerRequest("connect", echoCtx.Request())...,
		)
		traceID := span.SpanContext().TraceID()
		if traceID.IsValid() {
			shc.ResponseHeader().Set("X-Trace-ID", traceID.String())
		}
		return next(ctx, shc)
	}
}

// WrapUnary implements connect.Interceptor.
func (i *interceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, ar connect.AnyRequest) (connect.AnyResponse, error) {
		echoCtx, _ := http.EchoContextFromContext(ctx)
		span := trace.SpanFromContext(ctx)
		span.SetAttributes(
			httpconv.ServerRequest("connect", echoCtx.Request())...,
		)
		traceID := span.SpanContext().TraceID()
		res, err := next(ctx, ar)
		if traceID.IsValid() {
			res.Header().Set("X-Trace-ID", traceID.String())
		}
		return res, err
	}
}
