package service

import (
	"context"
	pb "double/api/admin/v1"
)

//func (s *AdminService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
//	return &pb.ListUserReply{}, nil
//}

func (s *AdminService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	//s.log.WithContext(ctx).Infof("%s service running", "admin")
	s.log.Debugf("%s service running", "admin")
	// 调用其他服务use case的方法
	list, _ := s.uc.List(ctx)
	res := make([]*pb.UserDetail, 0)
	for _,v := range list {
		res = append(res, &pb.UserDetail{
			Id:        v.ID,
			UserName:  v.UserName,
			ClubName:  v.ClubName,
			CreatedAt: v.CreatedAt,
		})
	}
	return &pb.ListUserReply{UserDetails: res}, nil
}
