package entity

import (
	"context"
	"gin-essential/contextx"

	"gorm.io/gorm"
)

func getDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	trans, ok := contextx.FromTrans(ctx)
	if ok && !contextx.FromNoTrans(ctx) {
		db, ok := trans.(*gorm.DB)
		if ok {
			if contextx.FromTransLock(ctx) {
				db = db.Set("gorm:query_option", "FOR UPDATE")
			}
			return db
		}
	}
	return defDB
}

// getDBWithModel ...
func getDBWithModel(ctx context.Context, defDB *gorm.DB, m interface{}) *gorm.DB {
	return getDB(ctx, defDB).Model(m)
}
