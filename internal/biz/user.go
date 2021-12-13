package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Hello string
}

type UserRepo interface {
	CreateUser(context.Context, *User) error
	UpdateUser(context.Context, *User) error
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Create(ctx context.Context, g *User) error {
	return uc.repo.CreateUser(ctx, g)
}

func (uc *UserUsecase) Update(ctx context.Context, g *User) error {
	return uc.repo.UpdateUser(ctx, g)
}
