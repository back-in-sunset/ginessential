package router

import (
	"gin-essential/router/api"
	"gin-essential/router/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
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
		auth.POST("msg", a.UserAPI.NatsMessage)
		auth.GET("click", a.UserAPI.QueryStatistics)
	}

	return app
}

// InitGinEngine 初始化gin引擎
func InitGinEngine(r IRouter) *gin.Engine {
	app := gin.New()

	// prefixes := r.Prefixes()

	// Router register
	r.Register(app)

	return app
}
