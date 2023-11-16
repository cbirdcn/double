package data

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	AccountStatusBanned = iota // 封号
	AccountStatusNormal // 正常状态的账号
)

// 时区
func CstSh () (*time.Location, error){
	return time.LoadLocation("Asia/Shanghai")
}

// 生成密码哈希（慢）
func PasswordHash(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// 校验密码哈希（慢）
func PasswordHashVerify(password string, hash string) (bool,error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err != nil, nil
}