package service

import (
	pb "double/api/admin/v1"
	"double/app/admin/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAdminService)

// 把所有其他rpc服务从xx.go中提出来，统一放到AdminService，方便注入

type AdminService struct {
	pb.UnimplementedAdminServer
	uc  *biz.UserUseCase
	// other rpc service
	log *log.Helper
}

func NewAdminService(uc *biz.UserUseCase, logger log.Logger) *AdminService {
	return &AdminService{uc: uc, log: log.NewHelper(logger)}
}

