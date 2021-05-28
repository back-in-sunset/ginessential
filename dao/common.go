package dao

import (
	"gin-essential/model/entity"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsn = "host=localhost user=postgres password=e.0369 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

// PostgresDB  PostgresDB
type PostgresDB struct {
	*gorm.DB
}

// PgDB 实例化
var PgDB PostgresDB

// InitDB mysql 初始化
func InitDB() {
	var once sync.Once

	once.Do(
		func() {
			db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err != nil {
				panic(err)
			}

			db.AutoMigrate(entity.User{})
			if os.Getenv("GOENV") == "dev" {
				log.Println("[INFO]> DB Starting.... IN Debug Mode ")
				db.Debug()
			}

			PgDB.DB = db
		},
	)

}
