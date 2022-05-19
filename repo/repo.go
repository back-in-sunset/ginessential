package repo

import (
	"gin-essential/repo/dao"

	demodao "gin-essential/repo/dao/demo"
	userdao "gin-essential/repo/dao/user"

	"github.com/google/wire"
)

// ModelSet model注入
var ModelSet = wire.NewSet(
	dao.TransSet,
	userdao.UserSet,
	demodao.DemoSet,
)
