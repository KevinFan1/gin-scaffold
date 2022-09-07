package middleware

import (
	"code/gin-scaffold/internal/global"
	"code/gin-scaffold/internal/vo"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// Error 自定义全局异常捕获
func Error() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				url := c.Request.URL
				method := c.Request.Method
				debug.PrintStack()
				global.Logger.Errorf("| url [%s] | method | [%s] | error [%s] | stack [%s] |", url, method, err, debug.Stack())

				switch t := err.(type) {
				case *vo.ApiError:
					c.JSON(t.Code, gin.H{
						"message": t.Message,
					})
				default:
					vo.ServerError(c)
				}
			}

			c.Abort()
		}()

		c.Next()
	}

}
