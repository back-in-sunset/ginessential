package model

import "gorm.io/gorm"

// User 用户
type User struct {
	gorm.Model
	Name      string `json:"name" gorm:"type:varchar(20);not null;"`
	Password  string `json:"password" gorm:"type:varchar(100);not null;"`
	Telephone string `json:"telephone" gorm:"type:varchar(110);not null;unique;"`
}
