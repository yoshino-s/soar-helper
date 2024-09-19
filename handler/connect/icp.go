package connect

import (
	"context"

	"connectrpc.com/connect"
	"gitlab.yoshino-s.xyz/yoshino-s/soar-helper/chinaz"
	v1 "gitlab.yoshino-s.xyz/yoshino-s/soar-helper/gen/v1"
	"gitlab.yoshino-s.xyz/yoshino-s/soar-helper/gen/v1/v1connect"
	"gitlab.yoshino-s.xyz/yoshino-s/soar-helper/persistent/db"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1connect.IcpQueryServiceHandler = (*IcpQueryService)(nil)

type IcpQueryService struct {
	chinaz *chinaz.Chinaz
	db     *db.Client
}

func NewIcpQueryService() *IcpQueryService {
	return &IcpQueryService{}
}

func (s *IcpQueryService) SetChinaz(chinaz *chinaz.Chinaz) {
	s.chinaz = chinaz
}

func (s *IcpQueryService) SetDB(db *db.Client) {
	s.db = db
}

func (s *IcpQueryService) Query(ctx context.Context, req *connect.Request[v1.QueryRequest]) (*connect.Response[v1.QueryResponse], error) {
	host := req.Msg.Host

	icp, cached, err := s.chinaz.Query(ctx, host)
	if err != nil {
		return nil, err
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
	res, cached, err := s.chinaz.BatchQuery(ctx, hosts)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	return connect.NewResponse(&v1.StatisticResponse{
		Total: int64(total),
	}), nil
}
