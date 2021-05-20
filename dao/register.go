package dao

import "gin-essential/model"

// IsTelePhoneExist 查询手机号是否存在
func (a *PostgresDB) IsTelePhoneExist(telephone string) bool {
	var user model.User
	a.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}

// Register 注册
func (a *PostgresDB) Register(newUser model.User) {
	a.Create(&newUser)
}
