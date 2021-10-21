package schema

import (
	"gin-essential/model/entity"
)

// Demo 用户
type Demo struct {
	DemoID int `json:"demo_id" gorm:"column:id"` // ID
	entity.DemoEntity
}

// DemoClickHose clickhouse data
type DemoClickHose struct {
	ID uint `json:"user_id" db:"id"`
	entity.DemoEntity
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
