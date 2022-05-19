package entity

// Role 角色
type Role struct {
	RoleName string `json:"role_name" gorm:"column:role_name;comment:角色名称"`
	MenuID   string `json:"menu_id" gorm:"menu_id;comment:菜单ID"`
}
