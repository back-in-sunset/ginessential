//go:build wireinject
// +build wireinject

package inject

import (
	"ginessential/logger"
	"ginessential/repo"
	"ginessential/repo/dao"
	"ginessential/router"
	"ginessential/router/api"
	"ginessential/srv"

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
