package entity

import (
	"context"
	"gin-essential/model/vo"
	"gin-essential/shared/id"

	"gorm.io/gorm"
)

// GetUserDB 获取用户存储
func GetUserDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return getDBWithTable(ctx, defDB, new(User))
}

// UserStauts .
type UserStauts int

const (
	// UserStautsOn 开启
	UserStautsOn UserStauts = 1
	// UserStautsOff 关闭
	UserStautsOff UserStauts = 2
)

// User 用户
type User struct {
	gorm.Model
	vo.User
	UserID id.UserID  `json:"user_id" gorm:"column:user_id;type:varchar(36);index;not null;"`
	Status UserStauts `json:"-" db:"status" gorm:"column:status;index;default:1;not null;"` // 状态(1:启用 2:停用)
}

// TableName 表名
func (*User) TableName() string {
	return "users"
}

// Users entity users
type Users []*User
