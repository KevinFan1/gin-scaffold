package router

import (
	"code/gin-scaffold/controller"
	"github.com/gin-gonic/gin"
)

func RecordsRouterInit(router *gin.RouterGroup) {
	group := router.Group("/records")
	{
		group.GET("/login", controller.LoginRecordListController)
	}
}
