package grpc

import (
	"context"

	"gitlab.yoshino-s.xyz/yoshino-s/icp-lookup/chinaz"
	v1 "gitlab.yoshino-s.xyz/yoshino-s/icp-lookup/gen/go/v1"
	"gitlab.yoshino-s.xyz/yoshino-s/icp-lookup/persistent/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IcpQueryService struct {
	v1.IcpQueryServiceServer
	chinaz *chinaz.Chinaz
	db     *db.Client
}

func New() *IcpQueryService {
	return &IcpQueryService{}
}

func (s *IcpQueryService) SetChinaz(chinaz *chinaz.Chinaz) {
	s.chinaz = chinaz
}

func (s *IcpQueryService) SetDB(db *db.Client) {
	s.db = db
}

func (s *IcpQueryService) Query(ctx context.Context, req *v1.QueryRequest) (*v1.QueryResponse, error) {
	host := req.Host

	icp, cached, err := s.chinaz.Query(ctx, host)
	if err != nil {
		return nil, err
	}

	return &v1.QueryResponse{
		Code:    uint32(codes.OK),
		Message: "ok",
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
		},
	}, nil
}

func (s *IcpQueryService) BatchQuery(ctx context.Context, req *v1.BatchQueryRequest) (*v1.BatchQueryResponse, error) {
	hosts := req.Hosts

	records := make([]*v1.IcpRecord, 0, len(hosts))
	for _, host := range hosts {
		icp, cached, err := s.chinaz.Query(ctx, host)
		if err != nil {
			continue
		}

		records = append(records, &v1.IcpRecord{
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
		})
	}

	return &v1.BatchQueryResponse{
		Code:    uint32(codes.OK),
		Message: "ok",
		Data:    records,
	}, nil
}

func (s *IcpQueryService) Statistic(ctx context.Context, _empty *emptypb.Empty) (*v1.StatisticResponse, error) {
	total, err := s.db.Icp.Query().Count(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.StatisticResponse{
		Total: int64(total),
	}, nil
}
