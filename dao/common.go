package dao

import (
	"gin-essential/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresDB  PostgresDB
type PostgresDB struct {
	*gorm.DB
}

// PgDB 实例化
var PgDB PostgresDB

// InitDB mysql 初始化
func InitDB() {
	dsn := "host=localhost user=postgres password=e.0369 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(model.User{})
	PgDB.DB = db
}
