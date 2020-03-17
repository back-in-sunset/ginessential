package model

import "github.com/jinzhu/gorm"

// User 用户
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null;"`
	Password  string `gorm:"type:varchar(100);not null;"`
	Telephone string `gorm:"type:varchar(110);not null;unique;"`
}
