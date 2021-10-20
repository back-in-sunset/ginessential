package schema

import (
	"gin-essential/model/entity"
	"gin-essential/pkg/errors"
	"gin-essential/pkg/util"
)

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
	UserName  string `form:"user_name"` // 用户名称
	Telephone string `form:"telephone"` // 手机号
}

// UserQueryResult 用户查询结果
type UserQueryResult struct {
	List       Users
	Pagination *PaginationResult
}

// Validate  数据验证
func (a *User) Validate() error {
	if len(a.Telephone) != 11 {
		return errors.New400Response("手机号格式不对")
	}

	if len(a.Password) < 6 {
		return errors.New400Response("密码少于6位")
	}

	if len(a.Name) == 0 {
		a.Name = util.RandomString(10)
	}
	return nil
}
