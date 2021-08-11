//+build wireinject

package inject

import (
	"gin-essential/dao"
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
		router.InitGinEngine,
		srv.SrvSet,
		api.APISet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
