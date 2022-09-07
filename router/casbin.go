package router

import (
	"code/gin-scaffold/controller"
	"code/gin-scaffold/internal/utils"
	"github.com/gin-gonic/gin"
)

func CasbinRouterInit(router *gin.RouterGroup) {
	casbinRouter := router.Group("/casbin")
	{
		casbinRouter.POST("", utils.LogDecorator(controller.CreateCasbinRuleController, "添加权限"))
		casbinRouter.DELETE("", controller.DeleteCasbinRuleController)
	}
}
