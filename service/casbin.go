package service

import (
	"code/gin-scaffold/internal/global"
	"code/gin-scaffold/models"
	"code/gin-scaffold/schemas"

	"github.com/gin-gonic/gin"
)

type CasbinSrv struct {
}

var CasbinService = new(CasbinSrv)

func (srv *CasbinSrv) Create(c *gin.Context) error {
	var rules schemas.CasbinChangeDto
	err := c.ShouldBindJSON(&rules)
	if err != nil {
		return err
	}
	_, err = global.Enforcer.AddPolicy(rules.Subject, rules.Object, rules.Action)
	if err != nil {
		return err
	}

	//添加redis缓存权限
	go models.CasbinCache{
		Code:   rules.Subject,
		Path:   rules.Object,
		Method: rules.Action,
	}.AddCasbinCache()
	return nil
}

func (srv *CasbinSrv) Delete(c *gin.Context) error {
	var rules schemas.CasbinChangeDto
	err := c.ShouldBindJSON(&rules)
	if err != nil {
		return err
	}

	_, err = global.Enforcer.RemovePolicy(rules.Subject, rules.Object, rules.Action)
	if err != nil {
		return err
	}

	//删除redis缓存权限
	go models.CasbinCache{
		Code:   rules.Subject,
		Path:   rules.Object,
		Method: rules.Action,
	}.DeleteCasbinCache()

	return nil
}
