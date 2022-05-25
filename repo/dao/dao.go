package dao

import (
	"context"
	"fmt"
	contextx "gin-essential/ctx"
	"gin-essential/model/entity"
	"gin-essential/schema"
	"log"
	"os"
	"time"

	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	pgdsn = "host=10.13.16.212 user=postgres password=e.0369 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
)

// Postgres postgres配置参数
type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// DSN 数据库连接串
func (a Postgres) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		a.Host, a.Port, a.User, a.DBName, a.Password, a.SSLMode)
}

// InitPgDB postgreSQL 初始化
func InitPgDB() *gorm.DB {
	// PostgresSQL 初始化
	pgDB, err := gorm.Open(postgres.Open(pgdsn),
		&gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := pgDB.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(63)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	if os.Getenv("GOENV") == "dev" {
		log.Println("[INFO]> DB Starting.... IN Debug Mode ")
		pgDB.Debug()
	}

	pgDB.AutoMigrate(&entity.User{}, &entity.Demo{})
	return pgDB
}

// WrapPageQuery 包装成带有分页的查询
func WrapPageQuery(db *gorm.DB, pp schema.PaginationParam, out interface{}) (*schema.PaginationResult, error) {
	if pp.OnlyCount {
		var count int64
		err := db.Count(&count).Error
		if err != nil {
			return nil, err
		}
		return &schema.PaginationResult{Total: int(count)}, nil
	} else if !pp.Pagination {
		err := db.Find(out).Error
		return nil, err
	}
	total, err := findPage(db, pp, out)
	if err != nil {
		return nil, err
	}
	return &schema.PaginationResult{
		Total:    total,
		Current:  pp.Current,
		PageSize: pp.PageSize,
	}, nil
}

func findPage(db *gorm.DB, pp schema.PaginationParam, out interface{}) (int, error) {
	var count int64
	err := db.Count(&count).Error
	if err != nil {
		return 0, err
	} else if count == 0 {
		return int(count), nil
	}
	current, pageSize := int(pp.Current), int(pp.PageSize)
	if current > 0 && pageSize > 0 {
		db = db.Offset((current - 1) * pageSize).Limit(pageSize)
	} else if pageSize > 0 {
		db = db.Limit(pageSize)
	}

	err = db.Find(out).Error
	return int(count), err
}

// FindOne 查询单条数据
func FindOne(ctx context.Context, db *gorm.DB, out interface{}) (bool, error) {
	db.First(out)
	if db.Error != nil {
		if db.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, db.Error
	}
	return true, nil
}

// TransSet 注入
var TransSet = wire.NewSet(wire.Struct(new(Trans), "*"))

// Trans 事务
type Trans struct {
	DB *gorm.DB
}

// Exec 事务执行
func (a *Trans) Exec(ctx context.Context, fn func(context.Context) error) error {
	if _, ok := contextx.FromTrans(ctx); ok {
		return fn(ctx)
	}

	return a.DB.Transaction(func(db *gorm.DB) error {
		return fn(contextx.NewTrans(ctx, db))
	})
}
