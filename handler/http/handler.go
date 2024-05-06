package http

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/yoshino-s/go-framework/application"
	"github.com/yoshino-s/go-framework/handlers/grpc_gateway"
	"github.com/yoshino-s/go-framework/handlers/http"
	v1 "gitlab.yoshino-s.xyz/yoshino-s/icp-lookup/gen/go/v1"
	"gitlab.yoshino-s.xyz/yoshino-s/icp-lookup/gen/openapiv2"
)

var _ application.Application = (*Handler)(nil)

type Handler struct {
	*http.Handler
	grpc v1.IcpQueryServiceServer
}

func New() *Handler {
	return &Handler{
		Handler: http.New(),
	}
}

func (h *Handler) SetGrpcServer(grpc v1.IcpQueryServiceServer) {
	h.grpc = grpc
}

func (h *Handler) Setup(ctx context.Context) {
	h.Handler.Setup(ctx)

	gh, err := grpc_gateway.New()
	if err != nil {
		panic(err)
	}

	v1.RegisterIcpQueryServiceHandlerServer(ctx, gh.ServeMux, h.grpc)

	h.Group("/api").Any("/*", echo.WrapHandler(gh))

	h.Swagger("/swagger", openapiv2.V1ServiceSwaggerJSON)
}
