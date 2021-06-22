package dao

import (
	"gin-essential/model/entity"
	"gin-essential/schema"
)

// UserPgDB 用户DB
type UserPgDB struct {
	*PostgresDB
}

// GetUserDB 获取UserDB
func GetUserDB() *UserPgDB {
	return &UserPgDB{
		&PgDB,
	}
}

// IsTelePhoneExist 查询手机号是否存在
func (a *UserPgDB) IsTelePhoneExist(telephone string) bool {
	var user entity.User
	a.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}

// Register 注册
func (a *UserPgDB) Register(user schema.User) error {
	a.Create(&entity.User{
		UserEntity: user.UserEntity,
	})
	if err := a.Error; err != nil {
		return err
	}
	return nil
}

// Query 查询
func (a UserPgDB) Query(params schema.UserQueryParams) (*schema.UserQueryResult, error) {
	db := a.DB
	if v := params.UserName; v != "" {
		db = db.Where("user_name like ?", "%"+v+"%")
	}
	var result schema.Users
	pr, err := WrapPageQuery(db, params.PaginationParam, &result)
	if err != nil {
		return &schema.UserQueryResult{}, err
	}
	return &schema.UserQueryResult{
		PageResult: pr,
		Data:       result,
	}, nil

}
