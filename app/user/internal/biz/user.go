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
	IsBanned bool
	CreatedAt string
}

// repo业务逻辑接口：在data中实现
type UserRepo interface {
	CreateUser(context.Context, *User) (int64, error)
	BanUser(context.Context, int64) (bool, error)
	UpdateUserName(context.Context, int64, string) (bool, error)
	UpdateUserPassword(context.Context, string, *User) (bool, error)
	GetUser(context.Context, int64) (*UserDetail, error)
	ListUser(context.Context) ([]*UserDetail, error)
	UserLogin(context.Context, *User) (*UserDetail, error)
	UserLogout(context.Context, *User) error
	UserIdCheckExist(context.Context, int64) bool
	UserNameCheckExist(context.Context, string) bool
	ClubNameCheckExist(context.Context, string) bool
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

func (uc *UserUseCase) Create(ctx context.Context, user *User) (int64, error) {
	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUseCase) Ban(ctx context.Context, id int64) (bool, error) {
	return uc.repo.BanUser(ctx, id)
}

func (uc *UserUseCase) UpdateUserName(ctx context.Context, id int64, newUserName string) (bool, error) {
	return uc.repo.UpdateUserName(ctx, id, newUserName)
}

func (uc *UserUseCase) UpdateUserPassword(ctx context.Context, newUserPassword string, user *User) (bool, error) {
	return uc.repo.UpdateUserPassword(ctx, newUserPassword, user)
}

func (uc *UserUseCase) Get(ctx context.Context, id int64) (*UserDetail, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UserUseCase) List(ctx context.Context) ([]*UserDetail, error) {
	return uc.repo.ListUser(ctx)
}

func (uc *UserUseCase) Login(ctx context.Context, user *User) (*UserDetail, error) {
	return uc.repo.UserLogin(ctx, user)
}

func (uc *UserUseCase) Logout(ctx context.Context, user *User) error {
	return uc.repo.UserLogout(ctx, user)
}