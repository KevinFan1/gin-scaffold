package router

import (
	"code/gin-scaffold/internal/settings"
	"code/gin-scaffold/internal/utils"
	"code/gin-scaffold/internal/vo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World! :)")
}

func Init(r *gin.Engine) {

	r.GET("/health", func(c *gin.Context) {
		vo.Ok(c, nil)
	})

	r.GET("/index", utils.LogDecorator(Index, "首页信息"))

	r.POST("/log", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	r.GET("/apis", func(c *gin.Context) {
		var data []map[string]string
		for _, r := range r.Routes() {
			data = append(data,
				map[string]string{
					"path":   r.Path,
					"method": r.Method,
				},
			)
		}
		vo.Ok(c, data)
	})

	prefix := settings.Setting.SystemBaseConfig.ApiPrefix

	v1Group := r.Group(prefix)
	v1Group.POST("/login", utils.JWTMiddleWareGenerator().LoginHandler)

	// 调用jwt和casbin中间件
	v1Group.Use(utils.JWTMiddleware())

	//添加需要校验权限的router
	UserRouterInit(v1Group)
	CasbinRouterInit(v1Group)
	RecordsRouterInit(v1Group)
}
