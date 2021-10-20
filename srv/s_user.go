package srv

import (
	"context"
	contextx "gin-essential/ctx"
	"gin-essential/dao"
	"gin-essential/logger"
	"gin-essential/schema"

	"github.com/google/wire"
)

// UserSet ..
var UserSet = wire.NewSet(wire.Struct(new(User), "*"))

// User ..
type User struct {
	UserDB *dao.User
	// UserChDB *dao.UserChDB
}

// IsTelePhoneExist 检查手机号是否存在
func (a *User) IsTelePhoneExist(ctx context.Context, telephone string) bool {
	result, err := a.UserDB.Query(ctx, schema.UserQueryParams{
		Telephone: telephone,
		PaginationParam: schema.PaginationParam{
			OnlyCount: true,
		}})
	if err != nil {
		return false
	}
	return result.Pagination.Total > 0
}

// Register 用户注册
func (a *User) Register(ctx context.Context, user schema.User) error {
	return a.UserDB.Create(ctx, user)
}

// QueryPage 查询分页数据
func (a *User) QueryPage(ctx context.Context, params schema.UserQueryParams) (*schema.UserQueryResult, error) {
	return a.UserDB.Query(ctx, params)
}

// Get 查询单条数据
func (a *User) Get(ctx context.Context, userID int) (*schema.User, error) {
	user, _ := contextx.FromUserID(ctx)

	logger.Logger.Info(user)
	return a.UserDB.Get(ctx, userID)
}

// Update 更新用户数据
func (a *User) Update(ctx context.Context, userID int, user schema.User) error {
	return a.UserDB.Update(ctx, userID, user)
}

// Delete 删除数据
func (a *User) Delete(ctx context.Context, userID int) error {
	return a.UserDB.Delete(ctx, userID)
}