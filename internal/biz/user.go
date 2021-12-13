package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

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
	CreatedAt string
}

// repo业务逻辑接口：在data中实现
type UserRepo interface {
	CreateUser(context.Context, *User) (int64, error)
	BanUser(context.Context, *User) error
	UpdateUser(context.Context, *User) error
	GetUser(context.Context, int64) (UserDetail, error)
	ListUser(context.Context, *User) error
	UserLogin(context.Context, *User) error
	UserLogout(context.Context, *User) error
	UserCheckExist(context.Context, *User) bool
	ClubCheckExist(context.Context, *User) bool
}

// useCase，在service中注入后，可以调用useCase的方法
type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

// useCase的方法：给service调用。内部调用repo的方法操作repo
func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUseCase) Create(ctx context.Context, user *User) (int64, error) {
	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUseCase) Ban(ctx context.Context, g *User) error {
	return uc.repo.BanUser(ctx, g)
}

func (uc *UserUseCase) Update(ctx context.Context, g *User) error {
	return uc.repo.UpdateUser(ctx, g)
}

func (uc *UserUseCase) Get(ctx context.Context, id int64) (UserDetail, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UserUseCase) List(ctx context.Context, g *User) error {
	return nil
}

func (uc *UserUseCase) Login(ctx context.Context, g *User) error {
	return nil
}

func (uc *UserUseCase) Logout(ctx context.Context, g *User) error {
	return nil
}