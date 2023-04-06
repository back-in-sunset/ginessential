package srv

import (
	usersrv "ginessential/srv/user"

	"github.com/google/wire"
)

// SrvSet srv注入
var SrvSet = wire.NewSet(
	usersrv.UserSet,
)
