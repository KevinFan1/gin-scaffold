package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CORS 跨域处理
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 计算总消耗时间
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "start_time", time.Now().UnixMilli()))

		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")

		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Header", "Content-Type,AccessToken,X-CSRF-TOKEN,Authorization,Token,X-TOKEN,X-USER,X-USER-ID")
		c.Header("Access-Control-Allow-Methods", "POST,PUT,OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Allow-Expose-Headers", "Content-Length,Access-Control-Allow-Origin")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
