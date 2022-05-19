package do

import (
	"context"
	"gin-essential/model/entity"

	"gorm.io/gorm"
)

// GetRoleDB 获取角色存储
func GetRoleDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return getDBWithTable(ctx, defDB, new(Role))
}

// Role 角色
type Role struct {
	gorm.Model
	entity.Role
	RoleID string `json:"role_id" gorm:"column:role_id;comment:角色ID"`
}

// TableName 表名
func (*Role) TableName() string {
	return "role"
}

// Roles Roles
type Roles []*Role
