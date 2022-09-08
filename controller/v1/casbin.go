package v1

import (
	"code/gin-scaffold/internal/vo"

	"github.com/gin-gonic/gin"
)

// CreateCasbinRule 创建casbin权限
func CreateCasbinRule(c *gin.Context) {
	err := casbinService.Create(c)
	if err != nil {
		vo.FailWithMsg(c, err.Error())
		return
	}
	vo.Ok(c, nil)
}

// DeleteCasbinRule 删除casbin权限
func DeleteCasbinRule(c *gin.Context) {
	err := casbinService.Delete(c)
	if err != nil {
		vo.FailWithMsg(c, err.Error())
		return
	}
	vo.Ok(c, nil)
}
