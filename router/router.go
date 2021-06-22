package router

import (
	"gin-essential/bll"
	"gin-essential/dao"
	"gin-essential/router/api"
	"gin-essential/router/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Init gin
func Init() http.Handler {
	e := gin.New()
	apis := initAPIs()
	e.Use(middleware.Cors())

	// e.Use()
	e.GET("heart_beat", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "ok"})
	})
	auth := e.Group("api/auth")
	{
		auth.POST("register", apis.Register)
		auth.POST("msg", apis.NatsMessage)
	}

	return e
}

// APIs apis
type APIs struct {
	api.User
}

func initAPIs() APIs {
	return APIs{
		api.User{BllUser: bll.User{UserDB: dao.GetUserDB()}},
	}

}
