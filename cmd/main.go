package main

import (
	"gin-essential/model/entity"
	"log"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
)

func main() {
	dsn := "tcp://localhost:9001?database=gorm&read_timeout=10&write_timeout=20&ebug=true"
	db, err := sqlx.Open("clickhouse", dsn)
	// db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	defer db.Close()
	var users []*entity.UserClickHose
	if err != nil {
		log.Printf("%v", err)
	}
	// db.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&entity.User{})
	// now := time.Now()

	err = db.Select(&users, "SELECT * EXCEPT (created_at, updated_at, deleted_at) FROM users WHERE id = ?", 0)
	if err != nil {
		log.Println(err)
	}
	for _, item := range users {
		log.Printf("%+v", item)
	}

}
