package connect

import (
	"bufio"
	"context"
	"io"
	"os"
	"os/exec"
	"strings"

	"connectrpc.com/connect"
	"github.com/go-errors/errors"
	"github.com/sourcegraph/conc"
	v1 "github.com/yoshino-s/soar-helper/internal/proto/v1"
	"github.com/yoshino-s/soar-helper/internal/proto/v1/v1connect"
)

var _ v1connect.RunnerServiceHandler = (*RunnerService)(nil)

type RunnerService struct {
}

func NewRunnerService() *RunnerService {
	return &RunnerService{}
}
func (s *RunnerService) Run(ctx context.Context, req *connect.Request[v1.RunRequest]) (*connect.Response[v1.RunResponse], error) {
	if len(req.Msg.Commands) == 0 {
		return nil, errors.Errorf("commands is empty")
	}
	cmd := exec.Command(req.Msg.Commands[0], req.Msg.Commands[1:]...)
	if req.Msg.Stdin != "" {
		cmd.Stdin = io.NopCloser(strings.NewReader(req.Msg.Stdin))
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, errors.New(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, errors.New(err)
	}

	code := 0

	err = cmd.Start()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			code = exitErr.ExitCode()
		} else {
			return nil, errors.New(err)
		}
	}

	stdoutContent, err := io.ReadAll(stdout)
	if err != nil {
		return nil, errors.New(err)
	}
	stderrContent, err := io.ReadAll(stderr)
	if err != nil {
		return nil, errors.New(err)
	}

	return connect.NewResponse(&v1.RunResponse{
		Stdout:   string(stdoutContent),
		Stderr:   string(stderrContent),
		ExitCode: int32(code),
	}), nil
}

func (s *RunnerService) RunStream(ctx context.Context, req *connect.Request[v1.RunRequest], stream *connect.ServerStream[v1.RunStreamData]) error {
	if len(req.Msg.Commands) == 0 {
		return errors.Errorf("commands is empty")
	}
	cmd := exec.Command(req.Msg.Commands[0], req.Msg.Commands[1:]...)
	if req.Msg.Stdin != "" {
		cmd.Stdin = io.NopCloser(strings.NewReader(req.Msg.Stdin))
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return errors.New(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return errors.New(err)
	}

	wg := conc.NewWaitGroup()

	wg.Go(func() {
		scanner := bufio.NewScanner(stdout)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			if err := stream.Send(&v1.RunStreamData{
				Type: v1.RunStreamDataType_STDOUT,
				Data: scanner.Text(),
			}); err != nil {
				panic(err)
			}
		}
	})

	wg.Go(func() {
		scanner := bufio.NewScanner(stderr)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			if err := stream.Send(&v1.RunStreamData{
				Type: v1.RunStreamDataType_STDERR,
				Data: scanner.Text(),
			}); err != nil {
				panic(err)
			}
		}
	})

	wg.Go(func() {
		code := 0
		err = cmd.Start()
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				code = exitErr.ExitCode()
			} else {
				panic(err)
			}
		}
		if err = stream.Send(&v1.RunStreamData{
			Type:     v1.RunStreamDataType_EXIT_CODE,
			ExitCode: int32(code),
		}); err != nil {
			panic(err)
		}
	})

	return wg.WaitAndRecover().AsError()
}

func (s *RunnerService) ReadFile(ctx context.Context, req *connect.Request[v1.ReadFileRequest]) (*connect.Response[v1.ReadFileResponse], error) {
	f, err := os.Open(req.Msg.Path)
	if err != nil {
		return nil, errors.New(err)
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		return nil, errors.New(err)
	}

	return connect.NewResponse(&v1.ReadFileResponse{
		Data: string(content),
	}), nil
}

func (s *RunnerService) WriteFile(ctx context.Context, req *connect.Request[v1.WriteFileRequest]) (*connect.Response[v1.WriteFileResponse], error) {
	f, err := os.OpenFile(req.Msg.Path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, errors.New(err)
	}
	defer f.Close()

	_, err = f.WriteString(req.Msg.Content)
	if err != nil {
		return nil, errors.New(err)
	}

	return connect.NewResponse(&v1.WriteFileResponse{}), nil
}
