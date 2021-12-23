package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// xxx:调用其他service：以user举例

// 数据模型
type User struct {
	ID int64
	UserName string
	Password string
	ClubName string
	IsBanned bool
	CreatedAt string
}

type UserDetail struct {
	ID int64
	UserName string
	ClubName string
	IsBanned bool
	CreatedAt string
}

type UserRepo interface {
	// 用listUser举例，调用服务的rpc方法
	ListUser(context.Context) ([]*UserDetail, error)
}

// useCase，在service中注入后，可以调用useCase的方法
type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

// useCase以及其方法：给service调用。内部调用repo的方法操作repo
func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUseCase) List(ctx context.Context) ([]*UserDetail, error) {
	return uc.repo.ListUser(ctx)
}
