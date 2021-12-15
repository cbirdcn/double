package data

import (
	"context"
	"double/api/user/v1"
	"double/internal/biz"
	"github.com/fatih/structs"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

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

func (r *userRepo) CreateUser(ctx context.Context, user *biz.User) (int64, error) {
	// todo:锁
	if r.UserNameCheckExist(ctx, user.UserName){
		return 0, user_v1.ErrorUserNameExist("user name %s exist", user.UserName)
	}
	if r.ClubNameCheckExist(ctx, user.ClubName){
		return 0, user_v1.ErrorClubNameExist("club name %s exist", user.ClubName)
	}
	// 忽略密码检查

	r.data.rdb.SAdd(ctx, getUserNameSetKey(), user.UserName)
	r.data.rdb.SAdd(ctx, getClubNameSetKey(), user.ClubName)
	// 获取自增id
	user.ID = getAutoIncrementId()

	// 密文存储password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, user_v1.ErrorUnknownError("save password error")
	}
	user.Password = string(hash)

	r.data.rdb.HMSet(ctx, getUserKey(user.ID), structs.Map(user))

	// 判断name->id的索引关系是否存在，不存在就设置，存在就退出
	// todo:锁
	if idFromName, _ := r.data.rdb.HGet(ctx, getIdByNameHashKey(), user.UserName).Int64(); idFromName != 0 {
		return 0, user_v1.ErrorUserNotFound("user exist: %s", user.UserName)
	} else {
		r.data.rdb.HSet(ctx, getIdByNameHashKey(), user.UserName, user.ID)
	}

	return user.ID, nil
}

func (r *userRepo) BanUser(ctx context.Context, id int64) (bool, error) {
	if !r.UserIdCheckExist(ctx, id) {
		return false, user_v1.ErrorUserNotFound("user id not found: %d", id)
	}

	values := make(map[string]interface{})
	values["IsBanned"] = true

	r.data.rdb.HSet(ctx, getUserKey(id), values)

	return true, nil
}

func (r *userRepo) UpdateUserName(ctx context.Context, id int64, newUserName string) (bool, error) {
	if !r.UserIdCheckExist(ctx, id) {
		return false, user_v1.ErrorUserNotFound("user id not found: %d", id)
	}

	// 新角色名已存在
	if r.UserNameCheckExist(ctx, newUserName) {
		return false, user_v1.ErrorUserNameExist("user name %s exist", newUserName)
	}

	// 旧角色和新角色名一致(可以省略)
	//if newUserName == r.data.rdb.HGet(ctx, getUserKey(id), "UserName").val() {
	//	return false, user_v1.ErrorUserNameSame("new user name %s same with current", newUserName)
	//}

	oldUserName := r.data.rdb.HGet(ctx, getUserKey(id), "UserName").Val()

	values := make(map[string]interface{})
	values["UserName"] = newUserName

	// todo:lock

	r.data.rdb.HSet(ctx, getUserKey(id), values)

	// 更新相关集合和索引中的昵称
	r.UserNameAdd(ctx, newUserName)
	r.UserNameDel(ctx, oldUserName)
	r.UserSetIdByChangedName(ctx, oldUserName, newUserName, id)

	return true, nil
}

func (r *userRepo) UpdateUserPassword(ctx context.Context, newUserPassword string, user *biz.User) (bool, error) {
	if !r.UserIdCheckExist(ctx, user.ID) {
		return false, user_v1.ErrorUserNotFound("user id not found: %d", user.ID)
	}

	// 旧密码正确
	if !r.UserComparePasswordSame(ctx, user.Password, user.ID) {
		return false, user_v1.ErrorUserPasswordDifferent("user old password wrong")
	}

	// 新密码不能是旧密码
	if r.UserComparePasswordSame(ctx, newUserPassword, user.ID) {
		return false, user_v1.ErrorUserPasswordExist("need different user password with now")
	}

	// 密文存储password
	hash, err := bcrypt.GenerateFromPassword([]byte(newUserPassword), bcrypt.DefaultCost)
	if err != nil {
		return false, user_v1.ErrorUnknownError("save password error")
	}
	values := make(map[string]interface{})
	values["Password"] = string(hash)

	r.data.rdb.HSet(ctx, getUserKey(user.ID), values)

	return true, nil
}

func (r *userRepo) GetUser(ctx context.Context, id int64) (biz.UserDetail, error) {
	values := r.data.rdb.HGetAll(ctx, getUserKey(id))
	val := values.Val()
	ID, _ := strconv.ParseInt(val["ID"], 10, 64)
	IsBanned, _ := strconv.ParseBool(val["IsBanned"])
	return biz.UserDetail{
		ID:        ID,
		UserName:  val["UserName"],
		ClubName:  val["ClubName"],
		CreatedAt: val["CreatedAt"],
		IsBanned: IsBanned,
	}, nil
}

//func (r *userRepo) ListUser(ctx context.Context) (users []*biz.User, error) {
//	// todo：目前还没有存储id list
//	return nil,nil
//}

func (r *userRepo) UserLogin(ctx context.Context, user *biz.User) (biz.UserDetail, error) {
	// todo:无挤号。无在线状态检测
	id := r.UserIdByName(ctx, user.UserName)
	if id == 0 {
		return biz.UserDetail{}, user_v1.ErrorUserNotFound("user not found: %s", user.UserName)
	}

	if !r.UserComparePasswordSame(ctx, user.Password, id) {
		return biz.UserDetail{}, user_v1.ErrorUserPasswordDifferent("user password wrong")
	}

	if res, err := r.GetUser(ctx, id); err != nil {
		return biz.UserDetail{}, user_v1.ErrorUserNotFound("user not found: %s", user.UserName)
	}else{
		// 没有被封禁
		if res.IsBanned {
			return biz.UserDetail{}, user_v1.ErrorUserBanned("user banned: %s", user.UserName)
		}else {
			return res, nil
		}
	}
}

func (r *userRepo) UserLogout(ctx context.Context, user *biz.User) error {
	panic("implement me")
}

// repo直接操作


func (r *userRepo) UserNameCheckExist(ctx context.Context, userName string) bool {
	// todo:锁
	get := r.data.rdb.SIsMember(ctx, getUserNameSetKey(), userName)
	return get.Val()
}

func (r *userRepo) UserNameDel(ctx context.Context, userName string) bool {
	// todo:锁
	if r.data.rdb.SRem(ctx, getUserNameSetKey(), userName).Val() == 1 {
		return true
	}
	return false
}

func (r *userRepo) UserNameAdd(ctx context.Context, userName string) bool {
	// todo:锁
	if r.data.rdb.SAdd(ctx, getUserNameSetKey(), userName).Val() == 1 {
		return true
	}
	return false
}

func (r *userRepo) ClubNameCheckExist(ctx context.Context, clubName string) bool {
	// todo:锁
	get := r.data.rdb.SIsMember(ctx, getClubNameSetKey(), clubName)
	return get.Val()
}

func (r *userRepo) UserIdCheckExist(ctx context.Context, id int64) bool {
	// todo:锁
	if r.data.rdb.Exists(ctx, getUserKey(id)).Val() == 1 {
		return true
	}
	return false
}

func (r *userRepo) UserComparePasswordSame(ctx context.Context, checkedPassword string, id int64) bool {

	// 密文存储
	hash := r.data.rdb.HGet(ctx, getUserKey(id), "Password").Val()

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(checkedPassword))

	if err == nil {
		return true
	}

	return false
}


// 略：密码有效性Check

func (r *userRepo) UserIdByName(ctx context.Context, userName string) int64 {
	// todo:hash field会很多，需要分key
	if id, err := r.data.rdb.HGet(ctx, getIdByNameHashKey(), userName).Int64(); err == nil {
		return id
	}
	return 0
}

func (r *userRepo) UserSetIdByChangedName(ctx context.Context, oldUserName string, newUserName string, id int64) int64 {
	// todo:hash field会很多，需要分key
	// todo:lock
	if res := r.data.rdb.HSet(ctx, getIdByNameHashKey(), newUserName, id).Val(); res == 1 {
		if res := r.data.rdb.HDel(ctx, getIdByNameHashKey(), oldUserName).Val(); res == 1 {
			return id
		} else {
			return 0
		}
	}
	return 0
}

