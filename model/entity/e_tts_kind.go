package entity

import (
	"context"

	"gorm.io/gorm"
)

// TTSTone tts音色
type TTSTone struct {
	gorm.Model
	TTSToneEntity
}

// GetTTSToneDB 获取tts tone存储
func GetTTSToneDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return getDBWithTable(ctx, defDB, new(TTSTone))
}

// TTSToneEntity ..
type TTSToneEntity struct {
	Name      string `json:"name" db:"name" gorm:"column:name;type:varchar(20);not null;"`
	Code      string `json:"code" db:"code" gorm:"column:code;type:varchar(11);"`
	TypeName  string `json:"type_name" db:"type_name" gorm:"column:type_name;type:varchar(11);"`
	SceneName string `json:"scene_name" db:"scene_name" gorm:"column:scene_name;size:11;"` //
}

// TableName ..
func (a TTSTone) TableName() string {
	return "m_tts_tone"
}
