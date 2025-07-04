package connect

import (
	"context"

	"connectrpc.com/connect"
	"github.com/go-errors/errors"
	"github.com/yoshino-s/entproto/runtime"
	"github.com/yoshino-s/go-framework/application"
	"github.com/yoshino-s/soar-helper/internal/beian"
	"github.com/yoshino-s/soar-helper/internal/persistent/db"
	"github.com/yoshino-s/soar-helper/internal/proto/entpb"
	"github.com/yoshino-s/soar-helper/internal/proto/entpb/entpbservice"
	v1 "github.com/yoshino-s/soar-helper/internal/proto/v1"
	"github.com/yoshino-s/soar-helper/internal/proto/v1/v1connect"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ v1connect.IcpQueryServiceHandler = (*IcpQueryServiceHandler)(nil)

type IcpQueryServiceHandler struct {
	*application.EmptyApplication
	Chinaz *beian.Beian `inject:""`
	DB     *db.Client   `inject:""`
}

func NewIcpQueryServiceHandler() *IcpQueryServiceHandler {
	return &IcpQueryServiceHandler{
		EmptyApplication: application.NewEmptyApplication("IcpQueryServiceHandler"),
	}
}

func (s *IcpQueryServiceHandler) Query(ctx context.Context, req *connect.Request[v1.QueryRequest]) (*connect.Response[entpb.Icp], error) {
	host := req.Msg.Host

	icp, _, err := s.Chinaz.Query(ctx, host, req.Msg.NoCache)
	if err != nil {
		return nil, errors.New(err)
	}

	return runtime.WrapResult(entpbservice.ToProtoIcp(icp))
}

func (s *IcpQueryServiceHandler) BatchQuery(ctx context.Context, req *connect.Request[v1.BatchQueryRequest]) (*connect.Response[v1.BatchQueryResponse], error) {
	hosts := req.Msg.Hosts

	res, _, errs, err := s.Chinaz.BatchQuery(ctx, hosts, req.Msg.NoCache)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	records, err := entpbservice.ToProtoIcpList(res)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&v1.BatchQueryResponse{
		Items:  records,
		Errors: errs,
	}), nil
}

func (s *IcpQueryServiceHandler) Statistic(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[v1.StatisticResponse], error) {
	total, err := s.DB.Icp.Query().Count(ctx)
	if err != nil {
		return nil, errors.New(err)
	}
	return connect.NewResponse(&v1.StatisticResponse{
		Total: int64(total),
	}), nil
}
