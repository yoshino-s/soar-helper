package s3

import (
	"context"
	"net/url"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/yoshino-s/go-framework/application"
	"github.com/yoshino-s/go-framework/configuration"
	"go.uber.org/zap"
)

type S3 struct {
	*application.EmptyApplication
	config config
	client *minio.Client
}

func New() *S3 {
	return &S3{
		EmptyApplication: application.NewEmptyApplication(),
		config:           config{},
	}
}

func (s *S3) Configuration() configuration.Configuration {
	return &s.config
}

func (s *S3) Setup(context.Context) {
	client, err := minio.New(s.config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(s.config.AccessKeyID, s.config.SecretAccessKey, ""),
		Secure: !s.config.Insecure,
		Region: s.config.Region,
	})
	if err != nil {
		panic(err)
	}
	s.client = client
}

func (s *S3) Client() *minio.Client {
	return s.client
}

func (s *S3) Upload(ctx context.Context, key string, path string, options minio.PutObjectOptions) (*url.URL, error) {
	info, err := s.client.FPutObject(ctx, s.config.Bucket, key, path, options)
	if err != nil {
		return nil, err
	}
	s.Logger.Debug("uploaded", zap.Any("info", info))

	return s.client.PresignedGetObject(ctx, s.config.Bucket, key, s.config.PresignedGetObjectExpires, url.Values{})
}
