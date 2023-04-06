package demodao

import (
	"context"
	"ginessential/model/entity"
	"ginessential/model/vo"
	"ginessential/repo/dao"
	"ginessential/schema"

	"github.com/google/wire"
	"gorm.io/gorm"
)

// DemoSet 注入Demo
var DemoSet = wire.NewSet(wire.Struct(new(Demo), "*"))

// Demo 用户DB
type Demo struct {
	// Logger
	PgDB *gorm.DB
}

// Create 创建
func (a *Demo) Create(ctx context.Context, demo schema.Demo) error {
	db := entity.GetDemoDB(ctx, a.PgDB).Create(&entity.Demo{
		DemoID: demo.DemoID,
		Demo:   demo.Demo,
	})
	if err := db.Error; err != nil {
		return err
	}
	return nil
}

// Query 查询
func (a *Demo) Query(ctx context.Context, params schema.DemoQueryParams) (*schema.DemoQueryResult, error) {
	db := entity.GetDemoDB(ctx, a.PgDB)

	if v := params.Name; v != "" {
		db = db.Where("name like ?", "%"+v+"%")
	}

	db = db.Order("id DESC")
	var result schema.Demos
	pr, err := dao.WrapPageQuery(db, params.PaginationParam, &result)
	if err != nil {
		return &schema.DemoQueryResult{}, err
	}
	return &schema.DemoQueryResult{
		Pagination: pr,
		List:       result,
	}, nil

}

// Get 获取单条数据
func (a *Demo) Get(ctx context.Context, demoID int) (*schema.Demo, error) {
	db := entity.GetDemoDB(ctx, a.PgDB)

	var demo schema.Demo
	_, err := dao.FindOne(ctx, db, &demo)
	if err != nil {
		return nil, err
	}

	return &demo, nil
}

// Update 更新
func (a *Demo) Update(ctx context.Context, userID int, user schema.Demo) error {
	db := entity.GetDemoDB(ctx, a.PgDB)

	db.Where("id = ?", userID).Updates(&user).Omit("id", "demo_id")
	if err := db.Error; err != nil {
		return err
	}
	return nil
}

// Delete 删除
func (a *Demo) Delete(ctx context.Context, demoID int) error {
	db := entity.GetDemoDB(ctx, a.PgDB).Where("demo_id = ?", demoID).Delete(vo.Demo{})
	if err := db.Error; err != nil {
		return err
	}
	return nil
}
