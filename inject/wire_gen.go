// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package inject

import (
	"gin-essential/bll"
	"gin-essential/dao"
	"gin-essential/router"
	"gin-essential/router/api"
)

// Injectors from wire.go:

// BuildInjector ..
func BuildInjector() (*Injector, func(), error) {
	db := dao.InitPgDB()
	user := &dao.User{
		PgDB: db,
	}
	chDB := dao.InitChDB()
	userChDB := &dao.UserChDB{
		ChDB: chDB,
	}
	bllUser := bll.User{
		UserDB:   user,
		UserChDB: userChDB,
	}
	apiUser := &api.User{
		UserBll: bllUser,
	}
	routerRouter := &router.Router{
		UserAPI: apiUser,
	}
	engine := router.InitGinEngine(routerRouter)
	injector := &Injector{
		Engine: engine,
	}
	return injector, func() {
	}, nil
}
