package dao

import "gin-essential/model/entity"

// IsTelePhoneExist 查询手机号是否存在
func (a *PostgresDB) IsTelePhoneExist(telephone string) bool {
	var user entity.User
	a.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}

// Register 注册
func (a *PostgresDB) Register(newUser entity.User) error {
	a.Create(&newUser)
	if err := a.Error; err != nil {
		return err
	}
	return nil
}
