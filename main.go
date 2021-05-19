package main

import (
	"gin-essential/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	dao.InitDB()

	http.ListenAndServe(":8080", engine)

	// r.POST("/api/auth/register", router.User.Register())
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
