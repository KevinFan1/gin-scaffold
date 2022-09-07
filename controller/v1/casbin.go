package v1

import (
	"code/gin-scaffold/internal/global"
	"code/gin-scaffold/internal/vo"
	"code/gin-scaffold/models"
	"code/gin-scaffold/schemas"
	"github.com/gin-gonic/gin"
)

// CreateCasbinRule 创建casbin权限
func CreateCasbinRule(c *gin.Context) {
	var rules schemas.CasbinChangeDto
	err := c.ShouldBindJSON(&rules)

	if err != nil {
		vo.FailWithMsg(c, err.Error())
		return
	}

	_, err = global.Enforcer.AddPolicy(rules.Subject, rules.Object, rules.Action)

	if err != nil {
		vo.FailWithMsg(c, err.Error())
		return
	}

	//添加redis缓存权限
	go models.CasbinCache{
		Code:   rules.Subject,
		Path:   rules.Object,
		Method: rules.Action,
	}.AddCasbinCache()
	vo.Ok(c, nil)
}

// DeleteCasbinRule 删除casbin权限
func DeleteCasbinRule(c *gin.Context) {
	var rules schemas.CasbinChangeDto
	err := c.ShouldBindJSON(&rules)
	if err != nil {
		vo.FailWithMsg(c, err.Error())
		return
	}

	_, err = global.Enforcer.RemovePolicy(rules.Subject, rules.Object, rules.Action)
	if err != nil {
		vo.FailWithMsg(c, err.Error())
		return
	}

	//删除redis缓存权限
	go models.CasbinCache{
		Code:   rules.Subject,
		Path:   rules.Object,
		Method: rules.Action,
	}.DeleteCasbinCache()
	vo.Ok(c, nil)
}
