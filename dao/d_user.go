package dao

import (
	"context"
	"gin-essential/model/entity"
	"gin-essential/schema"

	"github.com/google/wire"
	"gorm.io/gorm"
)

// UserSet 注入User
var UserSet = wire.NewSet(wire.Struct(new(User), "*"))

// User 用户DB
type User struct {
	PgDB *gorm.DB
}

// // UserChDB ..
// type UserChDB struct {
// 	ChDB *ChDB
// }

// IsTelePhoneExist 查询手机号是否存在
func (a *User) IsTelePhoneExist(ctx context.Context, telephone string) bool {
	var user entity.User

	db := entity.GetUserDB(ctx, a.PgDB)
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}

// Create 创建
func (a *User) Create(ctx context.Context, user schema.User) error {
	entity.GetUserDB(ctx, a.PgDB).Create(&entity.User{
		UserEntity: user.UserEntity,
	})
	if err := a.PgDB.Error; err != nil {
		return err
	}
	return nil
}

// Query 查询
func (a *User) Query(ctx context.Context, params schema.UserQueryParams) (*schema.UserQueryResult, error) {
	db := entity.GetUserDB(ctx, a.PgDB)

	if v := params.UserName; v != "" {
		db = db.Where("name like ?", "%"+v+"%")
	}

	db.Order("id DESC")
	var result schema.Users
	pr, err := WrapPageQuery(db, params.PaginationParam, &result)
	if err != nil {
		return &schema.UserQueryResult{}, err
	}
	return &schema.UserQueryResult{
		Pagination: pr,
		List:       result,
	}, nil

}

// Get 获取单条数据
func (a *User) Get(ctx context.Context, userID int) (*schema.User, error) {
	db := entity.GetUserDB(ctx, a.PgDB)

	var user schema.User
	db.Where("user_id = ?", userID).First(&user)
	if err := a.PgDB.Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Update 更新
func (a *User) Update(ctx context.Context, userID int, user schema.User) error {
	db := entity.GetUserDB(ctx, a.PgDB)

	db.Where("id = ?", userID).Updates(&user).Omit("id", "telephone", "email")
	if err := a.PgDB.Error; err != nil {
		return err
	}
	return nil
}

// // QueryStatistics 查询统计数据
// func (a *UserChDB) QueryStatistics(ctx context.Context) error {
// 	db := entity.GetUserDB(ctx, (*gorm.DB)(a.ChDB))

// 	var users schema.Users
// 	db.Find(&users)
// 	if err := db.Error; err != nil {
// 		log.Println(err)
// 	}
// 	log.Printf("%+v", users)
// 	return nil
// }

// // BatchCreate ..
// func (a *UserChDB) BatchCreate(ctx context.Context, users schema.Users) error {
// 	// db := entity.GetUserDB(ctx, (*gorm.DB)(a.ChDB))

// 	// log.Printf("%+v", users[0])
// 	// db.Create(*users[0])

// 	// if err := db.Error; err != nil {
// 	// 	log.Println(err)
// 	// 	return err
// 	// }
// 	return nil
// }
