package http

import (
	"context"

	"connectrpc.com/grpcreflect"
	"github.com/yoshino-s/go-framework/application"
	"github.com/yoshino-s/go-framework/common"
	"github.com/yoshino-s/go-framework/handlers/http"
	"gitlab.yoshino-s.xyz/yoshino-s/soar-helper/gen"
	"gitlab.yoshino-s.xyz/yoshino-s/soar-helper/gen/v1/v1connect"
	"go.akshayshah.org/connectproto"
	"google.golang.org/protobuf/encoding/protojson"
)

var _ application.Application = (*Handler)(nil)

type Handler struct {
	*http.Handler

	icpQueryHandler v1connect.IcpQueryServiceHandler
	runnerHandler   v1connect.RunnerServiceHandler
	toolsHandler    v1connect.ToolsServiceHandler
	s3Handler       v1connect.S3ServiceHandler
}

func New() *Handler {
	return &Handler{
		Handler: http.New(),
	}
}

func (h *Handler) SetIcpQueryHandler(handler v1connect.IcpQueryServiceHandler) {
	h.icpQueryHandler = handler
}

func (h *Handler) SetRunnerHandler(handler v1connect.RunnerServiceHandler) {
	h.runnerHandler = handler
}

func (h *Handler) SetToolsHandler(handler v1connect.ToolsServiceHandler) {
	h.toolsHandler = handler
}

func (h *Handler) SetS3Handler(handler v1connect.S3ServiceHandler) {
	h.s3Handler = handler
}

func (h *Handler) Setup(ctx context.Context) {
	common.MustNoNil(h.icpQueryHandler, h.runnerHandler, h.toolsHandler)
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

	opt := connectproto.WithJSON(
		protojson.MarshalOptions{EmitUnpopulated: true, EmitDefaultValues: true},
		protojson.UnmarshalOptions{DiscardUnknown: true},
	)

	h.HandleGrpc(v1connect.NewIcpQueryServiceHandler(h.icpQueryHandler, opt))
	h.HandleGrpc(v1connect.NewRunnerServiceHandler(h.runnerHandler, opt))
	h.HandleGrpc(v1connect.NewToolsServiceHandler(h.toolsHandler, opt))
	h.HandleGrpc(v1connect.NewS3ServiceHandler(h.s3Handler, opt))
}

func (h *Handler) Run(ctx context.Context) {
	h.Handler.Run(ctx)
}
