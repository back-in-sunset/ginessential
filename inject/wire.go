// +build wireinject

package inject

import (
	"gin-essential/dao"
	"gin-essential/logger"
	"gin-essential/router"
	"gin-essential/router/api"
	"gin-essential/srv"

	"github.com/google/wire"
)

// GenInjector ..
func GenInjector() (*Injector, func(), error) {
	wire.Build(
		dao.InitPgDB,
		dao.ModelSet,
		router.RouterSet,
		srv.SrvSet,
		api.APISet,
		logger.LoggerSet,
		router.GinSet,
		InjectorSet,
	)
	return new(Injector), func() { logger.Logger.Sync() }, nil
}
