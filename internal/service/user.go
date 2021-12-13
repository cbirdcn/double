package service

import (
	"context"
	"double/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"time"

	pb "double/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	id, err := s.uc.Create(ctx, &biz.User{
		UserName:  req.UserName,
		Password:  req.Password,
		ClubName:  req.ClubName,
		IsBanned:  false,
		CreatedAt: time.Now().Local().Format("2006-01-02 03:04:05"),
	})

	get, _ := s.uc.Get(ctx, id)

	userDetail := &pb.CreateUserReply{UserDetail: &pb.UserDetail{
		Id:       get.ID,
		UserName: get.UserName,
		ClubName: get.ClubName,
		CreatedAt: get.CreatedAt,
	}}
	return userDetail, err
}
func (s *UserService) UpdateUserName(ctx context.Context, req *pb.UpdateUserNameRequest) (*pb.UpdateUserNameReply, error) {
	return &pb.UpdateUserNameReply{}, nil
}
func (s *UserService) UpdateUserPassword(ctx context.Context, req *pb.UpdateUserPasswordRequest) (*pb.UpdateUserPasswordReply, error) {
	return &pb.UpdateUserPasswordReply{}, nil
}
func (s *UserService) BanUser(ctx context.Context, req *pb.BanUserRequest) (*pb.BanUserReply, error) {
	return &pb.BanUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
func (s *UserService) UserLogin(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginReply, error) {
	return &pb.UserLoginReply{}, nil
}
func (s *UserService) UserLogout(ctx context.Context, req *pb.UserLogoutRequest) (*pb.UserLogoutReply, error) {
	return &pb.UserLogoutReply{}, nil
}

//func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
//	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.GetName())
//
//	if in.GetName() == "error" {
//		return nil, v1.ErrorUserNotFound("user not found: %s", in.GetName())
//	}
//	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
//}
