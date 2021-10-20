package router

import (
	"gin-essential/router/api"
	"gin-essential/router/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	// swagger
	_ "gin-essential/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// RouterSet 注入router
var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

// GinSet gin
var GinSet = wire.NewSet(InitGinEngine)

// Router 路由管理器
type Router struct {
	UserAPI *api.User
}

// IRouter 注册路由
type IRouter interface {
	Registe(app *gin.Engine) error
	Prefixes() []string
}

// Registe 注册路由
func (a *Router) Registe(app *gin.Engine) error {
	a.RegisteAPI(app)
	return nil
}

// Prefixes 路由前缀列表
func (a *Router) Prefixes() []string {
	return []string{
		"/api/",
	}
}

// RegisteAPI 注册API
func (a *Router) RegisteAPI(app *gin.Engine) http.Handler {
	app.Use(middleware.Cors())
	// app.Group(strings.Join(a.Prefixes(app = ), ""))
	// e.Use()
	app.GET("heart_beat", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "ok"})
	})

	auth := app.Group("api/auth")
	{
		auth.GET("msg", a.UserAPI.NatsMessage)
	}
	users := app.Group("api/users")
	{
		users.GET("", a.UserAPI.Query)
		users.POST("", a.UserAPI.Register)
		users.GET(":id", a.UserAPI.Get)
		// users.GET("/:id/container/:tid", a.UserAPI.Get)
		users.GET(":id/start", a.UserAPI.Start)

	}

	return app
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name Jerry
// @contact.url http://www.swagger.io/supports
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2

// InitGinEngine 初始化gin引擎
func InitGinEngine(r IRouter) *gin.Engine {
	app := gin.New()

	// prefixes := r.Prefixes()

	// Router register
	r.Registe(app)

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return app
}
