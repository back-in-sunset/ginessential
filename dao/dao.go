package dao

import (
	"gin-essential/model/entity"
	"gin-essential/schema"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/clickhouse"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	pgdsn = "host=localhost user=postgres password=e.0369 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	chdsn = "tcp://localhost:9000?database=gorm&username=gorm&password=gorm&read_timeout=10&write_timeout=20"
)

// PostgresDB  PostgresDB
type PostgresDB struct {
	*gorm.DB
}

// ClickHouseDB clickhouse DB
type ClickHouseDB struct {
	*gorm.DB
}

// PgDB 实例化
var (
	PgDB PostgresDB
	ChDB ClickHouseDB
)

// InitDB mysql 初始化
func InitDB() {
	var once sync.Once

	once.Do(
		func() {
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
			PgDB.DB = pgDB

			// ClickHouse 初始化
			chDB, err := gorm.Open(clickhouse.Open(chdsn), &gorm.Config{})
			if err != nil {
				panic(err)
			}

			sqlDB, _ = chDB.DB()
			// SetMaxIdleConns 设置空闲连接池中连接的最大数量
			sqlDB.SetMaxIdleConns(10)
			// SetMaxOpenConns 设置打开数据库连接的最大数量。
			sqlDB.SetMaxOpenConns(100)
			// SetConnMaxLifetime 设置了连接可复用的最大时间。
			sqlDB.SetConnMaxLifetime(time.Hour)
			if os.Getenv("GOENV") == "dev" {
				log.Println("[INFO]> DB Starting.... IN Debug Mode ")
				chDB.Debug()
			}
			ChDB.DB = chDB
		},
	)

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
