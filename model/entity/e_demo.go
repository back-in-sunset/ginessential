package entity

import (
	"context"

	"gorm.io/gorm"
)

// GetDemoDB 获取用户存储
func GetDemoDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return getDBWithTable(ctx, defDB, new(Demo))
}

// Demo 用户
type Demo struct {
	gorm.Model
	DemoID string `json:"demo_id" gorm:"column:demo_id;index;type:varchar(36);not null;"` // demo id
	DemoEntity
}

// TableName 表名
func (*Demo) TableName() string {
	return "demo"
}

// DemoEntity 用户
type DemoEntity struct {
	Name      string  `json:"name" db:"name" gorm:"column:name;type:varchar(20);not null;"`
	Password  string  `json:"password" db:"password" gorm:"column:password;type:varchar(100);not null;"`
	Telephone string  `json:"telephone" db:"telephone" gorm:"column:telephone;type:varchar(110);not null;unique;"`
	Email     *string `json:"email" db:"email" gorm:"column:email;size:255;index;"`              // 邮箱
	Status    int     `json:"status" db:"status" gorm:"column:status;index;default:1;not null;"` // 状态(1:启用 2:停用)
}

// Demos demos
type Demos []*Demo
