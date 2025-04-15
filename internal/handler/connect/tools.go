package connect

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"

	"connectrpc.com/connect"
	"github.com/go-rod/rod"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/httpx/runner"
	"github.com/yoshino-s/go-framework/application"
	"github.com/yoshino-s/go-framework/errors"
	v1 "github.com/yoshino-s/soar-helper/internal/proto/v1"
	"github.com/yoshino-s/soar-helper/internal/proto/v1/v1connect"
	"github.com/yoshino-s/soar-helper/internal/s3"
	"github.com/yoshino-s/soar-helper/internal/screenshot"
	"github.com/yoshino-s/unauthor/scanner"
	"github.com/yoshino-s/unauthor/scanner/types"
	"go.uber.org/zap"
)

func init() {
	pageValue := reflect.ValueOf(rod.Page{})
	method := pageValue.MethodByName("Screenshot")
	if method.IsValid() && method.Kind() == reflect.Func {
		originalFunc := method.Interface().(func(bool) error)
		pageValue.MethodByName("Screenshot").Set(reflect.MakeFunc(method.Type(), func(args []reflect.Value) []reflect.Value {
			return reflect.ValueOf(originalFunc).Call([]reflect.Value{reflect.ValueOf(true)})
		}))
	}
}

var _ v1connect.ToolsServiceHandler = (*ToolsService)(nil)

var _ application.Application = (*ToolsService)(nil)

type ToolsService struct {
	*application.EmptyApplication
	s3 *s3.S3
}

func NewToolsService() *ToolsService {
	return &ToolsService{
		EmptyApplication: application.NewEmptyApplication("tools"),
	}
}

func (s *ToolsService) SetS3(s3 *s3.S3) {
	s.s3 = s3
}

func (t *ToolsService) Unauthor(ctx context.Context, req *connect.Request[v1.UnauthorRequest], stream *connect.ServerStream[v1.ExploitResult]) error {
	s := scanner.New()
	lock := &sync.Mutex{}

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

	cols := 80
	if req.Msg.ScreenshotWidth > 0 {
		cols = int(req.Msg.ScreenshotWidth)
	}

	takeScreenshot := func(result types.ScanFuncResult) (string, error) {
		tempFile, err := os.CreateTemp("", "screenshot-*.png")
		if err != nil {
			return "", err
		}
		defer tempFile.Close()

		if err := screenshot.CreateScreenshot([]string{
			result.Exploit,
		}, cols, result.Result, tempFile); err != nil {
			return "", err
		}

		if req.Msg.Upload && t.s3 != nil {
			if url, err := t.s3.Upload(ctx, fmt.Sprintf("cmd_screenshot/%s.png", uuid.NewString()), tempFile.Name(), minio.PutObjectOptions{}); err != nil {
				return "", err
			} else {
				return url.String(), nil
			}
		}

		return tempFile.Name(), nil
	}

	s.OutputFunc = func(result types.ScanFuncResult) {
		resp := &v1.ExploitResult{
			Target:  result.Target,
			Success: result.Success,
			Error:   result.Error,
			Result:  result.Result,
			Exploit: result.Exploit,
			Spent:   result.Time.Milliseconds(),
		}

		if req.Msg.Screenshot {
			if screenshot, err := takeScreenshot(result); err != nil {
				resp.Error = err.Error()
			} else {
				resp.Screenshot = screenshot
			}
		}

		lock.Lock()
		defer lock.Unlock()
		if err := stream.Send(resp); err != nil {
			t.Logger.Error("failed to send response", zap.Error(err))
		}
	}

	if f := config.ScanFuncs[config.Protocol]; f == nil {
		return connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("unsupported protocol: %s", config.Protocol))
	}

	s.Run(ctx)

	return nil
}

func (t *ToolsService) Httpx(ctx context.Context, req *connect.Request[v1.HttpxRequest], stream *connect.ServerStream[v1.ExploitResult]) error {
	lock := &sync.Mutex{}

	if req.Msg.Concurrent <= 0 {
		req.Msg.Concurrent = 50
	}
	if req.Msg.Timeout <= 0 {
		req.Msg.Timeout = int64(10 * time.Second)
	}

	if req.Msg.Upload && t.s3 == nil {
		return connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("s3 is not set"))
	}

	options := runner.Options{
		InputTargetHost: goflags.StringSlice(req.Msg.Targets),
		RequestURI:      req.Msg.Path,

		Threads: int(req.Msg.Concurrent),
		Timeout: int(time.Duration(req.Msg.Timeout) / time.Second),

		Screenshot:           req.Msg.Screenshot,
		UseInstalledChrome:   true,
		NoScreenshotFullPage: !req.Msg.FullScreenshot,
		NoScreenshotBytes:    true,
		NoHeadlessBody:       true,
		HeadlessOptionalArguments: []string{
			"--disable-features=HttpsUpgrades",
		},

		FollowRedirects:     req.Msg.FollowRedirects,
		FollowHostRedirects: req.Msg.FollowHostRedirects,

		OutputMatchStatusCode: req.Msg.MatchStatusCode,
		OutputMatchString:     []string{req.Msg.MatchString},

		NoFallbackScheme: true,

		StoreResponse:             true,
		StoreChain:                true,
		StoreResponseDir:          "/tmp/output",
		MaxResponseBodySizeToSave: math.MaxInt32,
		MaxResponseBodySizeToRead: math.MaxInt32,
		ScreenshotTimeout:         60 * time.Second,
		ScreenshotIdle:            1 * time.Second,
		Verbose:                   true,
		Delay:                     -1,
		HostMaxErrors:             30,
		MaxRedirects:              10,
		RandomAgent:               true,
		DisableUpdateCheck:        true,
		RateLimit:                 150,
		OnResult: func(r runner.Result) {
			// handle error
			if r.Err != nil {
				lock.Lock()
				defer lock.Unlock()

				stream.Send(&v1.ExploitResult{
					Target:  r.Input,
					Success: false,
					Error:   r.Err.Error(),
				})
			} else {
				screenshot := r.ScreenshotPath
				request := r.StoredResponsePath

				if req.Msg.Upload {
					if screenshot != "" {
						screenshotKey := strings.Join(strings.Split(screenshot, "/")[len(strings.Split(screenshot, "/"))-3:], "/")
						if url, err := t.s3.Upload(ctx, screenshotKey, screenshot, minio.PutObjectOptions{}); err != nil {
							t.Logger.Error("failed to upload screenshot", zap.Error(err))
							screenshot = ""
						} else {
							screenshot = url.String()
						}
					}

					if request != "" {
						requestKey := strings.Join(strings.Split(request, "/")[len(strings.Split(request, "/"))-3:], "/")
						if url, err := t.s3.Upload(ctx, requestKey, request, minio.PutObjectOptions{}); err != nil {
							t.Logger.Error("failed to upload request", zap.Error(err))
							request = ""
						} else {
							request = url.String()
						}
					}
				}

				r.ScreenshotPath = ""
				r.StoredResponsePath = ""

				lock.Lock()
				defer lock.Unlock()
				extra, _ := json.Marshal(map[string]interface{}{
					"request": request,
					"url":     r.URL,
					"status":  r.StatusCode,
				})
				stream.Send(&v1.ExploitResult{
					Target:  r.Input,
					Success: !r.Failed,
					Result:  r.JSON(nil),

					Screenshot: screenshot,
					Extra:      string(extra),
				})
			}
		},
	}

	if err := options.ValidateOptions(); err != nil {
		return errors.Wrap(err, "invalid options", http.StatusBadRequest)
	}

	httpxRunner, err := runner.New(&options)
	if err != nil {
		return errors.Wrap(err, "failed to create httpx runner")
	}
	defer httpxRunner.Close()

	httpxRunner.RunEnumeration()

	return nil
}
