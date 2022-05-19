package dao

import (
	"context"
	"gin-essential/model/do"
	"gin-essential/pkg/errors"
	"gin-essential/schema"

	"github.com/google/wire"
	"gorm.io/gorm"
)

// UserSet 注入User
var UserSet = wire.NewSet(wire.Struct(new(User), "*"))

// User 用户DB
type User struct {
	PgDB *gorm.DB
}

// Create 创建
func (a *User) Create(ctx context.Context, user schema.User) error {
	db := do.GetUserDB(ctx, a.PgDB).Create(&do.User{
		UserID: user.UserID,
		User:   user.User,
	})
	if err := db.Error; err != nil {
		return err
	}
	return nil
}

// Query 查询
func (a *User) Query(ctx context.Context, params schema.UserQueryParams) (*schema.UserQueryResult, error) {
	db := do.GetUserDB(ctx, a.PgDB)

	if v := params.UserName; v != "" {
		db = db.Where("name like ?", "%"+v+"%")
	}
	if v := params.Telephone; v != "" {
		db = db.Where("telephone = ?", v)
	}
	db.Order("id DESC")

	result := make(schema.Users, 0, params.PageSize)
	pr, err := WrapPageQuery(db, params.PaginationParam, &result)
	if err != nil {
		return &schema.UserQueryResult{}, err
	}
	return &schema.UserQueryResult{
		Pagination: pr,
		List:       result,
	}, nil

}

// Get 获取单条数据
func (a *User) Get(ctx context.Context, userID string) (*schema.User, error) {
	db := do.GetUserDB(ctx, a.PgDB).Where("user_id = ?", userID)

	var user schema.User
	ok, err := FindOne(ctx, db, &user)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.New500Response("new error")
	}

	return &user, nil
}

// Update 更新
func (a *User) Update(ctx context.Context, userID string, user schema.User) error {
	db := do.GetUserDB(ctx, a.PgDB)

	db.Where("id = ?", userID).Updates(&user).Omit("user_id", "telephone", "email")
	if err := db.Error; err != nil {
		return err
	}
	return nil
}

// Delete 删除
func (a *User) Delete(ctx context.Context, userID string) error {
	db := do.GetUserDB(ctx, a.PgDB).Where("user_id = ?", userID).Delete(&do.User{})
	if err := db.Error; err != nil {
		return err
	}
	return nil
}
