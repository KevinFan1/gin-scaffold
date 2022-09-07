package router

import (
	"code/gin-scaffold/controller"
	"github.com/gin-gonic/gin"
)

func UserRouterInit(router *gin.RouterGroup) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("", controller.UserListController)
		userGroup.POST("", controller.UserAdditionController)

		userGroup.GET("/detail", controller.UserDetailController)
		userGroup.GET("/info", controller.UserInfoController)

	}
}
