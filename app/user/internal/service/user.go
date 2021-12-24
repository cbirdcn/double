package service

import (
	"context"
	"double/app/user/internal/biz"
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
		CreatedAt: time.Now().Local().Format("2006-01-02 03:04:05"), // 服务器时区时间
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
	res, err := s.uc.UpdateUserName(ctx, req.Id, req.NewUserName)
	return &pb.UpdateUserNameReply{Res: res}, err
}
func (s *UserService) UpdateUserPassword(ctx context.Context, req *pb.UpdateUserPasswordRequest) (*pb.UpdateUserPasswordReply, error) {
	res, err := s.uc.UpdateUserPassword(ctx, req.NewUserPassword, &biz.User{
		ID:       req.Id,
		Password: req.OldUserPassword,
	})
	return &pb.UpdateUserPasswordReply{Res: res}, err
}
func (s *UserService) BanUser(ctx context.Context, req *pb.BanUserRequest) (*pb.BanUserReply, error) {
	res, err := s.uc.Ban(ctx, req.Id)
	return &pb.BanUserReply{Res: res}, err
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	get, _ := s.uc.Get(ctx, req.Id)

	userDetail := &pb.GetUserReply{UserDetail: &pb.UserDetail{
		Id:       get.ID,
		UserName: get.UserName,
		ClubName: get.ClubName,
		CreatedAt: get.CreatedAt,
	}}
	return userDetail, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	// 当前服务的日志
	s.log.WithContext(ctx).Infof("%s service running", "user")
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
func (s *UserService) UserLogin(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginReply, error) {
	detail, err := s.uc.Login(ctx, &biz.User{
		UserName: req.UserName,
		Password: req.Password,
	})

	if err == nil {
		return &pb.UserLoginReply{
			Res: true,
			UserDetail: &pb.UserDetail{
				Id:        detail.ID,
				UserName:  detail.UserName,
				ClubName:  detail.ClubName,
				CreatedAt: detail.CreatedAt,
			},
		}, err
	}

	return &pb.UserLoginReply{Res: false, UserDetail: &pb.UserDetail{}}, err
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
