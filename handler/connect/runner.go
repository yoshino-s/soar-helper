package connect

import (
	"bufio"
	"context"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"connectrpc.com/connect"
	"github.com/sourcegraph/conc"
	"github.com/yoshino-s/go-framework/errors"
	v1 "gitlab.yoshino-s.xyz/yoshino-s/soar-helper/gen/v1"
	"gitlab.yoshino-s.xyz/yoshino-s/soar-helper/gen/v1/v1connect"
)

var _ v1connect.RunnerServiceHandler = (*RunnerService)(nil)

type RunnerService struct {
}

func NewRunnerService() *RunnerService {
	return &RunnerService{}
}

func (s *RunnerService) Run(ctx context.Context, req *connect.Request[v1.RunRequest]) (*connect.Response[v1.RunResponse], error) {
	if len(req.Msg.Commands) == 0 {
		return nil, errors.New("commands is empty", http.StatusBadRequest)
	}
	cmd := exec.Command(req.Msg.Commands[0], req.Msg.Commands[1:]...)
	if req.Msg.Stdin != "" {
		cmd.Stdin = io.NopCloser(strings.NewReader(req.Msg.Stdin))
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	code := 0

	err = cmd.Start()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			code = exitErr.ExitCode()
		} else {
			return nil, err
		}
	}

	stdoutContent, err := io.ReadAll(stdout)
	if err != nil {
		return nil, err
	}
	stderrContent, err := io.ReadAll(stderr)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.RunResponse{
		Code:    0,
		Message: "success",
		Data: &v1.RunResult{
			Stdout:   string(stdoutContent),
			Stderr:   string(stderrContent),
			ExitCode: int32(code),
		},
	}), nil
}

func (s *RunnerService) RunStream(ctx context.Context, req *connect.Request[v1.RunRequest], stream *connect.ServerStream[v1.RunStreamData]) error {
	if len(req.Msg.Commands) == 0 {
		return errors.New("commands is empty", http.StatusBadRequest)
	}
	cmd := exec.Command(req.Msg.Commands[0], req.Msg.Commands[1:]...)
	if req.Msg.Stdin != "" {
		cmd.Stdin = io.NopCloser(strings.NewReader(req.Msg.Stdin))
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	wg := conc.NewWaitGroup()

	wg.Go(func() {
		scanner := bufio.NewScanner(stdout)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			if err := stream.Send(&v1.RunStreamData{
				Type: v1.RunStreamData_Type_STDOUT,
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
				Type: v1.RunStreamData_Type_STDERR,
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
			Type:     v1.RunStreamData_Type_EXIT_CODE,
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
		return nil, err
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.ReadFileResponse{
		Code:    0,
		Message: "success",
		Data:    string(content),
	}), nil
}

func (s *RunnerService) WriteFile(ctx context.Context, req *connect.Request[v1.WriteFileRequest]) (*connect.Response[v1.WriteFileResponse], error) {
	f, err := os.OpenFile(req.Msg.Path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, err = f.WriteString(req.Msg.Content)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.WriteFileResponse{
		Code:    0,
		Message: "success",
	}), nil
}
