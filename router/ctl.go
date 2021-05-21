package router

import (
	"gin-essential/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Init gin
func Init() http.Handler {
	e := gin.New()
	u := User{dao.PgDB}
	e.POST("api/auth/register", u.Register)
	e.POST("api/auth/msg", u.NatsMessage)
	return e
}
