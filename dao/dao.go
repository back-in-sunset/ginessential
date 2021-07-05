package dao

import (
	"gin-essential/model/entity"
	"gin-essential/schema"
	"log"
	"os"
	"time"

	"github.com/google/wire"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	pgdsn = "host=localhost user=postgres password=e.0369 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	chdsn = "tcp://localhost:9001?database=gorm&read_timeout=10&write_timeout=20"
)

// ModelSet model注入
var ModelSet = wire.NewSet(
	UserSet,
)

// InitPgDB postgreSQL 初始化
func InitPgDB() *gorm.DB {
	// PostgresSQL 初始化
	pgDB, err := gorm.Open(postgres.Open(pgdsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := pgDB.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	pgDB.AutoMigrate(entity.User{})
	if os.Getenv("GOENV") == "dev" {
		log.Println("[INFO]> DB Starting.... IN Debug Mode ")
		pgDB.Debug()
	}

	return pgDB
}

// InitChDB clickhouse 初始化
func InitChDB() *gorm.DB {
	// ClickHouse 初始化
	chDB, err := gorm.Open(clickhouse.Open(chdsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := chDB.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	// chDB.AutoMigrate(entity.User{})
	if os.Getenv("GOENV") == "dev" {
		log.Println("[INFO]> DB Starting.... IN Debug Mode ")
		chDB.Debug()
	}

	return chDB
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
		db.Offset((current - 1) * pageSize).Limit(pageSize)
	} else if pageSize > 0 {
		db = db.Limit(pageSize)
	}

	err = db.Find(out).Error
	return int(count), err
}
