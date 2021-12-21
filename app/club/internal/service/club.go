package service

import (
	"context"
	"double/app/club/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	pb "double/api/club/v1"
)

type ClubService struct {
	pb.UnimplementedClubServer
	uc  *biz.ClubUsecase
	log *log.Helper
}

func NewClubService(uc *biz.ClubUsecase, logger log.Logger) *ClubService {
	return &ClubService{uc: uc, log: log.NewHelper(logger)}
}

func (s *ClubService) CreateClub(ctx context.Context, req *pb.CreateClubRequest) (*pb.CreateClubReply, error) {
	return &pb.CreateClubReply{}, nil
}
func (s *ClubService) UpdateClub(ctx context.Context, req *pb.UpdateClubRequest) (*pb.UpdateClubReply, error) {
	return &pb.UpdateClubReply{}, nil
}
func (s *ClubService) DeleteClub(ctx context.Context, req *pb.DeleteClubRequest) (*pb.DeleteClubReply, error) {
	return &pb.DeleteClubReply{}, nil
}
func (s *ClubService) GetClub(ctx context.Context, req *pb.GetClubRequest) (*pb.GetClubReply, error) {
	return &pb.GetClubReply{}, nil
}
func (s *ClubService) ListClub(ctx context.Context, req *pb.ListClubRequest) (*pb.ListClubReply, error) {
	return &pb.ListClubReply{}, nil
}
