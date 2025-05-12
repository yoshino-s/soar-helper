package screenshot

import (
	"bytes"
	"context"
	"io"

	"go.opentelemetry.io/otel"
	oteltrace "go.opentelemetry.io/otel/trace"
)

func CreateScreenshot(
	ctx context.Context,
	command []string,
	cols int,
	result string,
	writer io.Writer,
) error {
	tracer := otel.GetTracerProvider().Tracer("github.com/yoshino-s/go-framework/internal/screenshot")
	ctx, span := tracer.Start(
		ctx,
		"CreateScreenshot",
		oteltrace.WithSpanKind(oteltrace.SpanKindInternal),
	)
	defer span.End()

	scaffold := NewImageCreator()

	scaffold.SetColumns(cols)
	scaffold.DrawShadow(true)
	scaffold.DrawDecorations(true)
	scaffold.ClipCanvas(false)

	scaffold.AddCommand(command...)

	scaffold.AddContent(bytes.NewReader([]byte(result)))
	scaffold.Write(writer)

	return nil
}
