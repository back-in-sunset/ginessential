package schema

import "gin-essential/model/entity"

// User 用户
type User struct {
	entity.UserEntity
}

// Users 用户列表
type Users []*User

// UserQueryParams 查询接口
type UserQueryParams struct {
}
