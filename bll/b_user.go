package bll

import (
	"context"
	"gin-essential/dao"
	"gin-essential/schema"
	"log"

	"github.com/google/wire"
)

// UserSet ..
var UserSet = wire.NewSet(wire.Struct(new(User), "*"))

// User ..
type User struct {
	UserDB   *dao.User
	UserChDB *dao.UserChDB
}

// IsTelePhoneExist 检查手机号是否存在
func (a *User) IsTelePhoneExist(ctx context.Context, telephone string) bool {
	return a.UserDB.IsTelePhoneExist(ctx, telephone)
}

// Register 用户注册
func (a *User) Register(ctx context.Context, user schema.User) error {
	return a.UserDB.Create(ctx, user)
}

// QueryPage 查询分页数据
func (a *User) QueryPage(ctx context.Context, params schema.UserQueryParams) (*schema.UserQueryResult, error) {
	return a.UserDB.Query(ctx, params)
}

// QueryStatistics ..
func (a *User) QueryStatistics(ctx context.Context) error {
	userResult, err := a.UserDB.Query(ctx, schema.UserQueryParams{})
	if err != nil {
		log.Println(err)
	}
	err = a.UserChDB.BatchCreate(ctx, userResult.Data)
	if err != nil {
		log.Println(err)
	}
	return a.UserChDB.QueryStatistics(ctx)
}
