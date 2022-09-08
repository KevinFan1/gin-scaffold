package router

import (
	"code/gin-scaffold/controller"
	"code/gin-scaffold/internal/utils"
	"github.com/gin-gonic/gin"
)

func UserRouterInit(router *gin.RouterGroup) {
	group := router.Group("/user")
	{
		group.GET("", utils.LogDecorator(controller.GetUserList, "用户列表"))
		group.POST("", controller.CreateUser)

		group.GET("/detail", controller.GetUserDetail)
		group.GET("/info", controller.GetUserInfo)

	}
}
