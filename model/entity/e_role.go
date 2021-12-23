package entity

import (
	"context"

	"gorm.io/gorm"
)

// GetRoleDB 获取角色存储
func GetRoleDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return getDBWithTable(ctx, defDB, new(Role))
}

// Role 角色
type Role struct {
	gorm.Model
	RoleEntity
	RoleID string `json:"role_id" gorm:"column:role_id;comment:角色ID"`
}

// TableName 表名
func (*Role) TableName() string {
	return "role"
}

// RoleEntity 角色
type RoleEntity struct {
	RoleName string `json:"role_name" gorm:"column:role_name;comment:角色名称"`
	MenuID   string `json:"menu_id" gorm:"menu_id;comment:菜单ID"`
}

// Roles Roles
type Roles []*Role
