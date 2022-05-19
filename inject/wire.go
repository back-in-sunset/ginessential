//go:build wireinject
// +build wireinject

package inject

import (
	"gin-essential/logger"
	"gin-essential/repo"
	"gin-essential/repo/dao"
	"gin-essential/router"
	"gin-essential/router/api"
	"gin-essential/srv"

	"github.com/google/wire"
)

// GenInjector ..
func GenInjector() (*Injector, func(), error) {
	wire.Build(
		dao.InitPgDB,
		repo.ModelSet,
		srv.SrvSet,
		api.APISet,
		router.RouterSet,
		logger.LoggerSet,
		router.GinSet,
		InjectorSet,
	)
	return new(Injector), func() { logger.Logger.Sync() }, nil
}
