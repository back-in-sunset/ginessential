package schema

import (
	"gin-essential/model/entity"
	"gin-essential/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
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
	UserName string `form:"user_name"` // 用户名称
}

// UserQueryResult 用户查询结果
type UserQueryResult struct {
	Data       Users
	PageResult *PaginationResult
}

// Validate  数据验证
func (a *User) Validate(c *gin.Context) bool {
	if len(a.Telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "must be 11 numbers"})
		return false
	}

	if len(a.Password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "must longer 6 numbers"})
		return false
	}

	if len(a.Name) == 0 {
		a.Name = util.RandomString(10)
	}
	return true
}
