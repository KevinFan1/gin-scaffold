package utils

import (
	"code/gin-scaffold/models"
	"context"
	"github.com/gin-gonic/gin"
)

func StringInSlice(s string, lst []string) bool {
	for _, ele := range lst {
		if s == ele {
			return true
		}
	}
	return false
}

// GetCurrentUser 获取当前request的用户
func GetCurrentUser(c *gin.Context) *models.User {
	data := c.Request.Context().Value("user")
	if data == nil {
		return nil
	}
	return data.(*models.User)
}

// LogDecorator 装饰器记录操作日志记录
func LogDecorator(f func(c *gin.Context), action string) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "log", action))
		f(c)
	}
}
