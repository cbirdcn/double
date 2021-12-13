package data

import (
	"context"
	"double/api/user/v1"
	"double/internal/biz"
	"fmt"
	"github.com/fatih/structs"
	"github.com/go-kratos/kratos/v2/log"
	"strconv"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (r *userRepo) UserCheckExist(ctx context.Context, user *biz.User) bool {
	// todo:锁
	get := r.data.rdb.SIsMember(ctx, getUserNameSetKey(), user.UserName)
	return get.Val()
}

func (r *userRepo) ClubCheckExist(ctx context.Context, user *biz.User) bool {
	// todo:锁
	get := r.data.rdb.SIsMember(ctx, getClubNameSetKey(), user.UserName)
	return get.Val()
}

// 略：密码有效性Check

func (r *userRepo) CreateUser(ctx context.Context, user *biz.User) (int64, error) {
	// todo:锁
	if r.UserCheckExist(ctx, user){
		return 0, user_v1.ErrorUserNameExist("user name %s exist", user.UserName)
	}
	if r.ClubCheckExist(ctx, user){
		return 0, user_v1.ErrorClubNameExist("club name %s exist", user.ClubName)
	}
	// 忽略密码检查

	r.data.rdb.SAdd(ctx, getUserNameSetKey(), user.UserName)
	r.data.rdb.SAdd(ctx, getClubNameSetKey(), user.ClubName)
	// 获取自增id
	user.ID = getAutoIncrementId()

	fmt.Println(user)

	r.data.rdb.HMSet(ctx, getUserKey(user.ID), structs.Map(user))

	return user.ID, nil
}

func (r *userRepo) BanUser(ctx context.Context, user *biz.User) error {
	panic("implement me")
}

func (r *userRepo) UpdateUser(ctx context.Context, g *biz.User) error {
	return nil
}

func (r *userRepo) GetUser(ctx context.Context, id int64) (biz.UserDetail, error) {
	values := r.data.rdb.HGetAll(ctx, getUserKey(id))
	val := values.Val()
	ID, _ := strconv.ParseInt(val["ID"], 10, 64)
	//CreatedAt, _ := time.ParseInLocation("2006-01-02 15:04:05", val["CreatedAt"], time.Local)
	return biz.UserDetail{
		ID:        ID,
		UserName:  val["UserName"],
		ClubName:  val["ClubName"],
		CreatedAt: val["CreatedAt"],
	}, nil
}

func (r *userRepo) ListUser(ctx context.Context, user *biz.User) error {
	panic("implement me")
}

func (r *userRepo) UserLogin(ctx context.Context, user *biz.User) error {
	panic("implement me")
}

func (r *userRepo) UserLogout(ctx context.Context, user *biz.User) error {
	panic("implement me")
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

