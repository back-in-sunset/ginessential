package entity

import (
	"context"

	"gorm.io/gorm"
)

// GetUserDB 获取用户存储
func GetUserDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return getDBWithTable(ctx, defDB, new(User))
}

// User 用户
type User struct {
	gorm.Model
	UserEntity
}

// TableName 表名
func (*User) TableName() string {
	return "users"
}

// UserEntity ..
type UserEntity struct {
	Name      string  `json:"name" db:"name" gorm:"column:name;type:varchar(20);not null;"`                        // 用户名
	Password  string  `json:"password" db:"password" gorm:"column:password;type:varchar(100);not null;"`           // 密码
	Telephone string  `json:"telephone" db:"telephone" gorm:"column:telephone;type:varchar(110);not null;unique;"` // 手机号
	Email     *string `json:"email" db:"email" gorm:"column:email;size:255;index;"`                                // 邮箱
	Status    int     `json:"status" db:"status" gorm:"column:status;index;default:1;not null;"`                   // 状态(1:启用 2:停用)
	UserID    string  `json:"user_id" gorm:"column:user_id;index;"`
}

// Users entity users
type Users []*User
