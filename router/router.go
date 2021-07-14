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

// Router 路由管理器
type Router struct {
	UserAPI *api.User
}

// IRouter 注册路由
type IRouter interface {
	Register(app *gin.Engine) error
	Prefixes() []string
}

// Register 注册路由
func (a *Router) Register(app *gin.Engine) error {
	a.RegisterAPI(app)
	return nil
}

// Prefixes 路由前缀列表
func (a *Router) Prefixes() []string {
	return []string{
		"/api/",
	}
}

// RegisterAPI 注册API
func (a *Router) RegisterAPI(app *gin.Engine) http.Handler {
	app.Use(middleware.Cors())
	// app.Group(strings.Join(a.Prefixes(app = ), ""))
	// e.Use()
	app.GET("heart_beat", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "ok"})
	})

	auth := app.Group("api/auth")
	{
		auth.POST("register", a.UserAPI.Register)
		auth.GET("msg", a.UserAPI.NatsMessage)
	}
	users := app.Group("api/users")
	{
		users.GET("", a.UserAPI.Query)
	}

	return app
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
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
	r.Register(app)

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return app
}
