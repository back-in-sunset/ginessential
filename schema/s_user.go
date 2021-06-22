package schema

import "gin-essential/model/entity"

// User 用户
type User struct {
	UserID int `json:"user_id" gorm:"column:id"` // 用户ID
	entity.UserEntity
}

// Users 用户列表
type Users []*User

// UserQueryParams 用户查询接口
type UserQueryParams struct {
	PaginationParam
	UserName string `form:"user_name"` // 用户名称
}

// UserQueryResult 用户查询结果
type UserQueryResult struct {
	Data       Users
	PageResult *PaginationResult
}
