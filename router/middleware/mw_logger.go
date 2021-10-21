package middleware

import (
	"fmt"
	"gin-essential/ginx"
	"gin-essential/logger"
	"mime"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ZapLogger ..
func ZapLogger(skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		p := c.Request.URL.Path
		method := c.Request.Method

		fields := logger.WithContext(c.Request.Context())

		start := time.Now()
		fields["ip"] = c.ClientIP()
		fields["method"] = method
		fields["url"] = c.Request.URL.String()
		fields["proto"] = c.Request.Proto
		fields["header"] = c.Request.Header
		fields["user_agent"] = c.GetHeader("User-Agent")
		fields["content_length"] = c.Request.ContentLength

		if method == http.MethodPost || method == http.MethodPut {
			mediaType, _, _ := mime.ParseMediaType(c.GetHeader("Content-Type"))
			if mediaType != "multipart/form-data" {
				if v, ok := c.Get(ginx.ReqBodyKey); ok {
					if b, ok := v.([]byte); ok {
						fields["body"] = string(b)
					}
				}
			}
		}
		c.Next()

		timeConsuming := time.Since(start).Nanoseconds() / 1e6
		status := c.Writer.Status()
		fields["res_status"] = status
		fields["res_length"] = c.Writer.Size()

		if v, ok := c.Get(ginx.ResBodyKey); ok {
			if b, ok := v.([]byte); ok {
				fields["res_body"] = string(b)
			}
		}
		if status == http.StatusOK {
			logger.Logger.Sugar().Info(
				fmt.Sprintf("[http] %s %s %s status:%d(%dms)", p, c.Request.Method, c.ClientIP(), c.Writer.Status(), timeConsuming),
				fields,
			)
		}

	}
}
