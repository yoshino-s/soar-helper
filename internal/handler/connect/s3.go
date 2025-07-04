package connect

import (
	"context"

	"connectrpc.com/connect"
	"github.com/go-errors/errors"
	"github.com/minio/minio-go/v7"
	"github.com/yoshino-s/go-framework/application"
	v1 "github.com/yoshino-s/soar-helper/internal/proto/v1"
	"github.com/yoshino-s/soar-helper/internal/proto/v1/v1connect"
	"github.com/yoshino-s/soar-helper/internal/s3"
)

var _ v1connect.S3ServiceHandler = (*S3ServiceHandler)(nil)

type S3ServiceHandler struct {
	*application.EmptyApplication
	S3 *s3.S3 `inject:""`
}

func NewS3ServiceHandler() *S3ServiceHandler {
	return &S3ServiceHandler{
		EmptyApplication: application.NewEmptyApplication("S3ServiceHandler"),
	}
}

func (s *S3ServiceHandler) Upload(ctx context.Context, req *connect.Request[v1.UploadRequest]) (*connect.Response[v1.UploadResponse], error) {
	url, err := s.S3.Upload(ctx, req.Msg.Key, req.Msg.Path, minio.PutObjectOptions{})
	if err != nil {
		return nil, errors.New(err)
	}
	return connect.NewResponse(&v1.UploadResponse{
		Url: url.String(),
	}), nil
}
