package entity

import (
	"context"
	"gin-essential/model/vo"
	"gin-essential/shared/id"

	"gorm.io/gorm"
)

// GetDemoDB 获取用户存储
func GetDemoDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return getDBWithTable(ctx, defDB, new(Demo))
}

// DemoStatus demo status
type DemoStatus int

const (
	// DemoStatusOn on
	DemoStatusOn DemoStatus = 1
	// DemoStatusOff off
	DemoStatusOff DemoStatus = 2
)

// Demo 用户
type Demo struct {
	gorm.Model
	vo.Demo
	DemoID id.DemoID  `json:"-" gorm:"column:demo_id;index;type:varchar(36);not null;"` // demo id
	Status DemoStatus `json:"-" gorm:"column:status;index;default:1;not null;"`         // 状态(1:启用 2:停用)
}

// Demos demos
type Demos []*Demo

// TableName 表名
func (*Demo) TableName() string {
	return "demos"
}
