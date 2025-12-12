package http

import (
	"context"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"connectrpc.com/otelconnect"
	"github.com/yoshino-s/go-framework/application"
	"github.com/yoshino-s/go-framework/handlers/http"
	connect_handler "github.com/yoshino-s/soar-helper/internal/handler/connect"
	gen "github.com/yoshino-s/soar-helper/internal/proto"
	"github.com/yoshino-s/soar-helper/internal/proto/v1/v1connect"
	"go.akshayshah.org/connectproto"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
)

var _ application.Application = (*Handler)(nil)

type Handler struct {
	*http.Handler
}

func New() *Handler {
	return &Handler{
		Handler: http.New(),
	}
}

func (h *Handler) Set(
	runnerServiceHandler *connect_handler.RunnerServiceHandler,
	toolsServiceHandler *connect_handler.ToolsServiceHandler,
	s3ServiceHandler *connect_handler.S3ServiceHandler,
) {
	reflector := grpcreflect.NewStaticReflector(
		v1connect.RunnerServiceName,
		v1connect.ToolsServiceName,
		v1connect.S3ServiceName,
	)

	h.HandleGrpc(grpcreflect.NewHandlerV1(reflector))
	h.HandleGrpc(grpcreflect.NewHandlerV1Alpha(reflector))

	var opts []connect.HandlerOption

	opts = append(opts, connectproto.WithJSON(
		protojson.MarshalOptions{EmitUnpopulated: true, EmitDefaultValues: true},
		protojson.UnmarshalOptions{DiscardUnknown: true},
	))

	otelInterceptor, err := otelconnect.NewInterceptor(otelconnect.WithTrustRemote())
	if err != nil {
		h.EmptyApplication.Logger.Fatal("failed to create otel interceptor", zap.Error(err))
	} else {
		opts = append(opts, connect.WithInterceptors(otelInterceptor, &interceptor{}))
	}

	h.HandleGrpc(v1connect.NewRunnerServiceHandler(runnerServiceHandler, opts...))
	h.HandleGrpc(v1connect.NewToolsServiceHandler(toolsServiceHandler, opts...))
	h.HandleGrpc(v1connect.NewS3ServiceHandler(s3ServiceHandler, opts...))
}

func (h *Handler) Setup(ctx context.Context) {
	h.Handler.Setup(ctx)

	h.Swagger("/swagger", gen.OpenAPI)
}

func (h *Handler) Run(ctx context.Context) {
	h.Handler.Run(ctx)
}
