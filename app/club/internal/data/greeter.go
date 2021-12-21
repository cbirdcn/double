package data

import (
	"context"
	"double/app/club/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type clubRepo struct {
	data *Data
	log  *log.Helper
}

// NewClubRepo .
func NewClubRepo(data *Data, logger log.Logger) biz.ClubRepo {
	return &clubRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *clubRepo) CreateClub(ctx context.Context, g *biz.Club) error {
	return nil
}

func (r *clubRepo) UpdateClub(ctx context.Context, g *biz.Club) error {
	return nil
}
