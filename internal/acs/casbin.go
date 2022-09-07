package acs

import (
	"code/gin-scaffold/internal/global"
	"code/gin-scaffold/models"
	"context"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/go-redis/redis/v8"
	"log"
)

func Setup() {

	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(global.DB, &models.CasbinRule{}, "t_casbin_rule")
	if err != nil {
		log.Panicf("加载adapter失败,%v\n", err)
	}

	global.Enforcer, err = casbin.NewEnforcer("internal/acs/rbac_model.conf", adapter)
	if err != nil {
		log.Panicf("加载rbac model失败,%v\n", err)
	}

	err = global.Enforcer.LoadPolicy()
	if err != nil {
		log.Panicf("加载rbac policy失败,%v\n", err)
	}

	//开启日志记录
	global.Enforcer.EnableLog(true)

}

func CheckPermission(code, path, method string) bool {

	cacheKey := models.CasbinRuleCacheKey
	redisClient := global.RedisClient
	ctx := context.Background()

	fieldKey := models.CasbinRuleFieldKey(code, path, method)
	val, err := redisClient.HGet(ctx, cacheKey, fieldKey).Result()

	if err == redis.Nil {
		result, err := global.Enforcer.Enforce(code, path, method)
		if err != nil {
			return false
		}
		redisClient.HSet(ctx, cacheKey, fieldKey, result)
		return result
	} else if err != nil {
		return false
	} else {
		if val == "1" {
			return true
		}
		return false
	}
}
