package do

import (
	"context"
	"gin-essential/model/entity"

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
	entity.Demo
}

// Demos demos
type Demos []*Demo

// TableName 表名
func (*Demo) TableName() string {
	return "demo"
}
