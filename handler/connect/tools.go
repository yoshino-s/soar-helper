package connect

import (
	"context"
	"fmt"
	"math"
	"strings"
	"time"

	"connectrpc.com/connect"
	"github.com/minio/minio-go/v7"
	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/httpx/runner"
	"github.com/yoshino-s/unauthor/scanner"
	"github.com/yoshino-s/unauthor/scanner/types"
	v1 "gitlab.yoshino-s.xyz/yoshino-s/soar-helper/gen/v1"
	"gitlab.yoshino-s.xyz/yoshino-s/soar-helper/gen/v1/v1connect"
	"gitlab.yoshino-s.xyz/yoshino-s/soar-helper/s3"
	"go.uber.org/zap"
)

var _ v1connect.ToolsServiceHandler = (*ToolsService)(nil)

type ToolsService struct {
	s3 *s3.S3
}

func NewToolsService() *ToolsService {
	return &ToolsService{}
}

func (s *ToolsService) SetS3(s3 *s3.S3) {
	s.s3 = s3
}

func (*ToolsService) Unauthor(ctx context.Context, req *connect.Request[v1.UnauthorRequest], stream *connect.ServerStream[v1.UnauthorResponse]) error {
	s := scanner.New()
	config, ok := s.Configuration().(*scanner.ScannerConfig)
	if !ok {
		return fmt.Errorf("unexpected configuration type: %T", s.Configuration())
	}
	config.Concurrent = int(req.Msg.Concurrent)
	if config.Concurrent <= 0 {
		config.Concurrent = 50
	}
	config.Timeout = time.Duration(req.Msg.Timeout)
	if config.Timeout <= 0 {
		config.Timeout = 5 * time.Second
	}
	config.Targets = req.Msg.Targets
	config.Protocol = req.Msg.Protocol

	s.OutputFunc = func(result types.ScanFuncResult) {
		resp := &v1.UnauthorResponse{
			Target:  result.Target,
			Success: result.Success,
			Error:   result.Error,
			Result:  result.Result,
		}
		stream.Send(resp)
	}

	if f := config.ScanFuncs[config.Protocol]; f == nil {
		return connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("unsupported protocol: %s", config.Protocol))
	}

	s.Run(ctx)

	return nil
}

func (t *ToolsService) Httpx(ctx context.Context, req *connect.Request[v1.HttpxRequest], stream *connect.ServerStream[v1.HttpxResponse]) error {

	if req.Msg.Concurrent <= 0 {
		req.Msg.Concurrent = 50
	}
	if req.Msg.Timeout <= 0 {
		req.Msg.Timeout = int64(10 * time.Second)
	}
	options := runner.Options{
		InputTargetHost:     goflags.StringSlice(req.Msg.Targets),
		Threads:             int(req.Msg.Concurrent),
		Timeout:             int(time.Duration(req.Msg.Timeout) / time.Second),
		Screenshot:          req.Msg.Screenshot,
		FollowRedirects:     req.Msg.FollowRedirects,
		FollowHostRedirects: req.Msg.FollowHostRedirects,

		StoreResponse:             true,
		StoreChain:                true,
		StoreResponseDir:          "output",
		MaxResponseBodySizeToSave: math.MaxInt32,
		MaxResponseBodySizeToRead: math.MaxInt32,
		ScreenshotTimeout:         60,
		Verbose:                   true,
		Delay:                     -1,
		HostMaxErrors:             30,
		MaxRedirects:              10,
		RandomAgent:               true,
		DisableUpdateCheck:        true,
		RateLimit:                 150,

		NoScreenshotBytes: true,
		NoHeadlessBody:    true,
		OnResult: func(r runner.Result) {
			// handle error
			if r.Err != nil {
				stream.Send(&v1.HttpxResponse{
					Target:  r.Input,
					Success: false,
					Error:   r.Err.Error(),
				})
			} else {
				screenshot := r.ScreenshotPath
				request := r.StoredResponsePath

				if req.Msg.Upload && t.s3 != nil {
					if screenshot != "" {
						screenshotKey := strings.Join(strings.Split(screenshot, "/")[len(strings.Split(screenshot, "/"))-3:], "/")
						if url, err := t.s3.Upload(ctx, screenshotKey, screenshot, minio.PutObjectOptions{}); err != nil {
							zap.L().Error("failed to upload screenshot", zap.Error(err))
						} else {
							screenshot = url.String()
						}
					}

					if request != "" {
						requestKey := strings.Join(strings.Split(request, "/")[len(strings.Split(request, "/"))-3:], "/")
						if url, err := t.s3.Upload(ctx, requestKey, request, minio.PutObjectOptions{}); err != nil {
							zap.L().Error("failed to upload request", zap.Error(err))
						} else {
							request = url.String()
						}
					}
				}

				r.ScreenshotPath = ""
				r.StoredResponsePath = ""

				stream.Send(&v1.HttpxResponse{
					Target:  r.Input,
					Success: !r.Failed,
					Result:  r.JSON(nil),

					Screenshot: screenshot,
					Request:    request,

					Url:        r.URL,
					StatusCode: int32(r.StatusCode),
				})
			}
		},
	}

	if err := options.ValidateOptions(); err != nil {
		return err
	}

	httpxRunner, err := runner.New(&options)
	if err != nil {
		return err
	}
	defer httpxRunner.Close()

	httpxRunner.RunEnumeration()

	return nil
}
