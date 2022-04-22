package do

import (
	"context"
	"gin-essential/model/entity"

	"gorm.io/gorm"
)

// GetUserDB 获取用户存储
func GetUserDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return getDBWithTable(ctx, defDB, new(User))
}

// User 用户
type User struct {
	gorm.Model
	UserID string `json:"user_id" gorm:"column:user_id;type:varchar(36);index;not null;"`
	entity.User
}

// TableName 表名
func (*User) TableName() string {
	return "users"
}

// Users entity users
type Users []*User
