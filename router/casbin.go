package router

import (
	"code/gin-scaffold/controller"
	"code/gin-scaffold/internal/utils"
	"github.com/gin-gonic/gin"
)

func CasbinRouterInit(router *gin.RouterGroup) {
	group := router.Group("/casbin")
	{
		group.POST("", utils.LogDecorator(controller.CreateCasbinRuleController, "添加权限"))
		group.DELETE("", controller.DeleteCasbinRuleController)
	}
}
