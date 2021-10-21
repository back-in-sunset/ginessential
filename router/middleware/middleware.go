package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SkipperFunc 定义中间件跳过函数
type SkipperFunc func(*gin.Context) bool

// Cors 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		// 可将将* 替换为指定的域名
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

// JWT jwt
// func JWT() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var (
// 			code int
// 			data interface{}
// 		)
// 		code = e.SUCCESS

// 		c.Next()
// 	}
// }

// SkipHandler 忽略handler
func SkipHandler(c *gin.Context, skipperFuncs ...SkipperFunc) bool {
	for _, skipperFunc := range skipperFuncs {
		if skipperFunc(c) {
			return true
		}
	}
	return false
}
