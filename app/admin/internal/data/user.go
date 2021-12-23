package data

import (
	"context"
	userv1 "double/api/user/v1"
	"double/app/admin/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// todo:page limit
func (r *userRepo) ListUser(ctx context.Context) ([]*biz.UserDetail, error) {
	// fmt.Println(r.data.uc)
	// 需要用uc去访问user Server的ListUser。请注意newData、Data、Discover整个过程中都要有uc。
	// 如果注入不完全，或endpoint错误，此时r.data.uc = nil。并触发panic
	reply, err := r.data.uc.ListUser(ctx, &userv1.ListUserRequest{})
	if err != nil {
		return nil, err
	}

	ret := make([]*biz.UserDetail, 0)
	for _,v := range reply.UserDetails {
		ret = append(ret, &biz.UserDetail{
			ID:        v.Id,
			UserName:  v.UserName,
			ClubName:  v.ClubName,
			CreatedAt: v.CreatedAt,
		})
	}

	return ret,nil
}
