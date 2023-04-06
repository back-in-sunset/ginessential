package repo

import (
	"ginessential/repo/dao"

	demodao "ginessential/repo/dao/demo"
	userdao "ginessential/repo/dao/user"

	"github.com/google/wire"
)

// ModelSet model注入
var ModelSet = wire.NewSet(
	dao.TransSet,
	userdao.UserSet,
	demodao.DemoSet,
)
