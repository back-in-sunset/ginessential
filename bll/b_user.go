package bll

import (
	"gin-essential/dao"
	"gin-essential/schema"
)

// User ..
type User struct {
	UserDB dao.PostgresDB
}

// IsTelePhoneExist 检查手机号是否存在
func (a *User) IsTelePhoneExist(telephone string) bool {
	return a.UserDB.IsTelePhoneExist(telephone)
}

// Register 用户注册
func (a *User) Register(user schema.User) error {
	a.UserDB.Register(user)
	return nil
}
