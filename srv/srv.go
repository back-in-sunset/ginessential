package srv

import "github.com/google/wire"

// SrvSet srv注入
var SrvSet = wire.NewSet(
	UserSet,
)
