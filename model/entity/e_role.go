package entity

import (
	"context"
	"gin-essential/model/do"
	"gin-essential/shared/id"

	"gorm.io/gorm"
)

// GetRoleDB 获取角色存储
func GetRoleDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return getDBWithTable(ctx, defDB, new(Role))
}

// Role 角色
type Role struct {
	gorm.Model
	do.Role
	RoleID id.RoleID `json:"role_id" gorm:"column:role_id;comment:角色ID"`
}

// TableName 表名
func (*Role) TableName() string {
	return "role"
}

// Roles Roles
type Roles []*Role
