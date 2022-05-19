package do

import (
	"context"

	contextx "gin-essential/ctx"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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
func getDBWithTable(ctx context.Context, defDB *gorm.DB, tabler schema.Tabler) *gorm.DB {
	return getDB(ctx, defDB).Table(tabler.TableName())
}
