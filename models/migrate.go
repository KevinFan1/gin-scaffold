package models

import (
	"code/gin-scaffold/internal/global"
)

func AutoMigrate() {
	err := global.DB.AutoMigrate(
		&Role{}, &User{}, &LoginRecord{}, &OperationRecord{},
	)
	if err != nil {
		global.Logger.Panicf("自动迁移失败,err: %v", err)
	}
	global.Logger.Info("自动迁移成功.")
}
