package router

import (
	"code/gin-scaffold/controller"
	"github.com/gin-gonic/gin"
)

func UserRouterInit(router *gin.RouterGroup) {
	group := router.Group("/user")
	{
		group.GET("", controller.UserListController)
		group.POST("", controller.UserAdditionController)

		group.GET("/detail", controller.UserDetailController)
		group.GET("/info", controller.UserInfoController)

	}
}
