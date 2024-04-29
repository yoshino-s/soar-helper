package grpc

import (
	"context"

	v1 "github.com/yoshino-s/go-template/gen/go/v1"
)

type GoTemplateServiceService struct {
	v1.GoTemplateServiceServer
}

func New() *GoTemplateServiceService {
	return &GoTemplateServiceService{}
}

func (GoTemplateServiceService) Echo(_ context.Context, msg *v1.StringMessage) (*v1.StringMessage, error) {
	return &v1.StringMessage{Value: msg.GetValue()}, nil
}
