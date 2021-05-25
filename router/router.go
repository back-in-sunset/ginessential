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
	e.POST("api/auth/register", u.Register)
	e.POST("api/auth/msg", u.NatsMessage)
	return e
}
