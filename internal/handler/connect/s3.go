package connect

import (
	"context"

	"connectrpc.com/connect"
	"github.com/go-errors/errors"
	"github.com/minio/minio-go/v7"
	v1 "github.com/yoshino-s/soar-helper/internal/proto/v1"
	"github.com/yoshino-s/soar-helper/internal/proto/v1/v1connect"
	"github.com/yoshino-s/soar-helper/internal/s3"
)

var _ v1connect.S3ServiceHandler = (*S3Service)(nil)

type S3Service struct {
	s3 *s3.S3
}

func NewS3Service() *S3Service {
	return &S3Service{}
}

func (s *S3Service) SetS3(s3 *s3.S3) {
	s.s3 = s3
}

func (s *S3Service) Upload(ctx context.Context, req *connect.Request[v1.UploadRequest]) (*connect.Response[v1.UploadResponse], error) {
	url, err := s.s3.Upload(ctx, req.Msg.Key, req.Msg.Path, minio.PutObjectOptions{})
	if err != nil {
		return nil, errors.New(err)
	}
	return connect.NewResponse(&v1.UploadResponse{
		Url: url.String(),
	}), nil
}
