package http

import (
	"context"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"connectrpc.com/otelconnect"
	"github.com/yoshino-s/go-framework/application"
	"github.com/yoshino-s/go-framework/handlers/http"
	"github.com/yoshino-s/go-framework/log"
	"github.com/yoshino-s/go-framework/utils"
	gen "github.com/yoshino-s/soar-helper/internal/proto"
	"github.com/yoshino-s/soar-helper/internal/proto/v1/v1connect"
	"go.akshayshah.org/connectproto"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
)

var _ application.Application = (*Handler)(nil)

type Handler struct {
	*http.Handler

	icpQueryHandler v1connect.IcpQueryServiceHandler `inject:""`
	runnerHandler   v1connect.RunnerServiceHandler   `inject:""`
	toolsHandler    v1connect.ToolsServiceHandler    `inject:""`
	s3Handler       v1connect.S3ServiceHandler       `inject:""`
}

func New() *Handler {
	return &Handler{
		Handler: http.New(),
	}
}

func (h *Handler) Setup(ctx context.Context) {
	utils.MustNoNil(h.icpQueryHandler, h.runnerHandler, h.toolsHandler)
	h.Handler.Setup(ctx)

	h.Swagger("/swagger", gen.OpenAPI)

	reflector := grpcreflect.NewStaticReflector(
		v1connect.IcpQueryServiceName,
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
		h.EmptyApplication.Logger.Fatal("failed to create otel interceptor", zap.Error(err), log.Context(ctx))
	} else {
		opts = append(opts, connect.WithInterceptors(otelInterceptor, &interceptor{}))
	}

	h.HandleGrpc(v1connect.NewIcpQueryServiceHandler(h.icpQueryHandler, opts...))
	h.HandleGrpc(v1connect.NewRunnerServiceHandler(h.runnerHandler, opts...))
	h.HandleGrpc(v1connect.NewToolsServiceHandler(h.toolsHandler, opts...))
	h.HandleGrpc(v1connect.NewS3ServiceHandler(h.s3Handler, opts...))
}

func (h *Handler) Run(ctx context.Context) {
	h.Handler.Run(ctx)
}
