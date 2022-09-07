package models

import (
	"code/gin-scaffold/internal/global"
	"context"
	"fmt"
)

var CasbinRuleCacheKey = "rbac:permission"

func CasbinRuleFieldKey(code, path, method string) string {
	return fmt.Sprintf("%v_%v_%v", code, path, method)
}

type CasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:100;uniqueIndex:unique_index"`
	V0    string `gorm:"size:100;uniqueIndex:unique_index"`
	V1    string `gorm:"size:100;uniqueIndex:unique_index"`
	V2    string `gorm:"size:100;uniqueIndex:unique_index"`
	V3    string `gorm:"size:100;uniqueIndex:unique_index"`
	V4    string `gorm:"size:100;uniqueIndex:unique_index"`
	V5    string `gorm:"size:100;uniqueIndex:unique_index"`
}

type CasbinCache struct {
	Code   string
	Path   string
	Method string
}

func (c CasbinCache) DeleteCasbinCache() {
	global.RedisClient.HDel(context.Background(), CasbinRuleCacheKey, CasbinRuleFieldKey(c.Code, c.Path, c.Method))
}

func (c CasbinCache) AddCasbinCache() {
	global.RedisClient.HSet(context.Background(), CasbinRuleCacheKey, CasbinRuleFieldKey(c.Code, c.Path, c.Method), "1")
}
