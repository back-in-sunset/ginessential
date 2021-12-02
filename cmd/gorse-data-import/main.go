package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

}

// UserActivity ..
type UserActivity struct {
}

// Item ..
type Item struct {
	ItemID string `gorm:"column:item_id"`
}

const dsn = "payfun_dms_service:Paymentunion123@@tcp(rm-8vbjy34g96075qpoklo.mysql.zhangbei.rds.aliyuncs.com:3408)/dms?charset=utf8mb4&parseTime=True&loc=Local"

var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

// QueryDevice ..
func QueryDevice() {
	db.Raw(`SELECT a.serialnum as item_id, a.activated_coun_name, a.activated_city_name, a.activated_province_name, b.label_map_id, c.name
	FROM aiot_device_attribute a 
	LEFT JOIN tdm_label_maps_sn b ON a.serialnum = b.sn
	LEFT JOIN tdm_label_maps c ON b.label_map_id = c.name
	LIMIT 1000`)
}
