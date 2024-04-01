package grpc_handler

import (
	"context"

	v1 "github.com/yoshino-s/go-template/gen/go/v1"
)

type GoTemplateServiceService struct {
	v1.GoTemplateServiceServer
}

func NewGoTemplateServiceService() (*GoTemplateServiceService, error) {
	return &GoTemplateServiceService{}, nil
}

func (GoTemplateServiceService) Echo(_ context.Context, msg *v1.StringMessage) (*v1.StringMessage, error) {
	return &v1.StringMessage{Value: msg.GetValue()}, nil
}
