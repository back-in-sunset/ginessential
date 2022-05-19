package schema

import (
	"gin-essential/model/do"
	"gin-essential/shared/id"
)

// Demo 用户
type Demo struct {
	DemoID id.DemoID `json:"demo_id" gorm:"column:demo_id"` // ID
	do.Demo
}

// Demos 用户列表
type Demos []*Demo

// DemoQueryParams 用户查询接口
type DemoQueryParams struct {
	PaginationParam
	Name string `form:"name"` // 名称
}

// DemoQueryResult 用户查询结果
type DemoQueryResult struct {
	List       Demos
	Pagination *PaginationResult
}
