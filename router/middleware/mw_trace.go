package middleware

import (
	contextx "gin-essential/ctx"
	"gin-essential/pkg/utils"

	"github.com/gin-gonic/gin"
)

// TraceMiddleware 增加traceID
func TraceMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		traceID := c.GetHeader("X-Request-Id")
		if traceID == "" {
			traceID = utils.NewTraceID()
		}

		ctx := contextx.NewTraceID(c.Request.Context(), traceID)
		c.Request = c.Request.WithContext(ctx)
		c.Writer.Header().Set("X-Trace-Id", traceID)

		c.Next()
	}
}
