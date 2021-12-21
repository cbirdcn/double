package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Club struct {
	Hello string
}

type ClubRepo interface {
	CreateClub(context.Context, *Club) error
	UpdateClub(context.Context, *Club) error
}

type ClubUsecase struct {
	repo ClubRepo
	log  *log.Helper
}

func NewClubUsecase(repo ClubRepo, logger log.Logger) *ClubUsecase {
	return &ClubUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ClubUsecase) Create(ctx context.Context, g *Club) error {
	return uc.repo.CreateClub(ctx, g)
}

func (uc *ClubUsecase) Update(ctx context.Context, g *Club) error {
	return uc.repo.UpdateClub(ctx, g)
}
