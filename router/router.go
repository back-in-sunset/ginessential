package router

import (
	"gin-essential/dao"
	"gin-essential/router/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Init gin
func Init() http.Handler {
	e := gin.New()
	u := api.User{UserDB: dao.PgDB}
	// e.Use()
	e.GET("heart_beat", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "ok"})
	})
	auth := e.Group("api/auth")
	{
		auth.POST("register", u.Register)
		auth.POST("msg", u.NatsMessage)
	}

	return e
}
