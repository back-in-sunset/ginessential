package entity

// Demo 用户
type Demo struct {
	Name      string  `json:"name" db:"name" gorm:"column:name;type:varchar(20);not null;"`
	Password  string  `json:"password" db:"password" gorm:"column:password;type:varchar(100);not null;"`
	Telephone string  `json:"telephone" db:"telephone" gorm:"column:telephone;type:varchar(110);not null;unique;"`
	Email     *string `json:"email" db:"email" gorm:"column:email;size:255;index;"`              // 邮箱
	Status    int     `json:"status" db:"status" gorm:"column:status;index;default:1;not null;"` // 状态(1:启用 2:停用)
}
