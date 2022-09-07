package middleware

import (
	"code/gin-scaffold/internal/settings"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// IpLimiter ip限制中间件
func IpLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIp := c.ClientIP()
		allowAccess := false

		for _, ip := range settings.Setting.SystemBaseConfig.AllowedHost {
			if ip == "*" || clientIp == ip {
				allowAccess = true
				break
			}
		}

		if !allowAccess {
			c.String(http.StatusForbidden, fmt.Sprintf("Your Ip: %v not in ip permission list.", clientIp))
			c.Abort()
		}

		c.Next()
	}

}
