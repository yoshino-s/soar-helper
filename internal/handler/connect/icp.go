package connect

import (
	"context"

	"connectrpc.com/connect"
	"github.com/yoshino-s/go-framework/errors"
	"github.com/yoshino-s/soar-helper/internal/beian"
	"github.com/yoshino-s/soar-helper/internal/persistent/db"
	v1 "github.com/yoshino-s/soar-helper/internal/proto/v1"
	"github.com/yoshino-s/soar-helper/internal/proto/v1/v1connect"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1connect.IcpQueryServiceHandler = (*IcpQueryService)(nil)

type IcpQueryService struct {
	chinaz *beian.Beian
	db     *db.Client
}

func NewIcpQueryService() *IcpQueryService {
	return &IcpQueryService{}
}

func (s *IcpQueryService) SetChinaz(chinaz *beian.Beian) {
	s.chinaz = chinaz
}

func (s *IcpQueryService) SetDB(db *db.Client) {
	s.db = db
}

func (s *IcpQueryService) Query(ctx context.Context, req *connect.Request[v1.QueryRequest]) (*connect.Response[v1.QueryResponse], error) {
	host := req.Msg.Host

	icp, cached, err := s.chinaz.Query(ctx, host, req.Msg.NoCache)
	if err != nil {
		return nil, errors.Wrap(err, "query chinaz icp error")
	}

	return connect.NewResponse(&v1.QueryResponse{
		Data: &v1.IcpRecord{
			Id:        int64(icp.ID),
			Host:      icp.Host,
			City:      icp.City,
			Province:  icp.Province,
			Company:   icp.Company,
			Owner:     icp.Owner,
			Type:      icp.Type,
			Homepage:  icp.Homepage,
			Permit:    icp.Permit,
			CreatedAt: timestamppb.New(icp.CreatedAt),
			UpdatedAt: timestamppb.New(icp.UpdatedAt),
			Cached:    cached,
			Input:     host,
			WebName:   icp.WebName,
		},
	}), nil
}

func (s *IcpQueryService) BatchQuery(ctx context.Context, req *connect.Request[v1.BatchQueryRequest]) (*connect.Response[v1.BatchQueryResponse], error) {
	hosts := req.Msg.Hosts

	records := make([]*v1.IcpRecord, len(hosts))
	res, cached, err := s.chinaz.BatchQuery(ctx, hosts, req.Msg.NoCache)
	if err != nil {
		return nil, errors.Wrap(err, "query chinaz icp error")
	}
	for i, icp := range res {
		records[i] = &v1.IcpRecord{
			Id:        int64(icp.ID),
			Host:      icp.Host,
			City:      icp.City,
			Province:  icp.Province,
			Company:   icp.Company,
			Owner:     icp.Owner,
			Type:      icp.Type,
			Homepage:  icp.Homepage,
			Permit:    icp.Permit,
			CreatedAt: timestamppb.New(icp.CreatedAt),
			UpdatedAt: timestamppb.New(icp.UpdatedAt),
			Cached:    cached[i],
			Input:     hosts[i],
			WebName:   icp.WebName,
		}
	}

	return connect.NewResponse(&v1.BatchQueryResponse{
		Data: records,
	}), nil
}

func (s *IcpQueryService) Statistic(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[v1.StatisticResponse], error) {
	total, err := s.db.Icp.Query().Count(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "query icp count error")
	}
	return connect.NewResponse(&v1.StatisticResponse{
		Total: int64(total),
	}), nil
}
