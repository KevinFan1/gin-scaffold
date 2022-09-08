package main

import (
	"code/gin-scaffold/internal/acs"
	"code/gin-scaffold/internal/database"
	"code/gin-scaffold/internal/global"
	"code/gin-scaffold/internal/gredis"
	"code/gin-scaffold/internal/logging"
	"code/gin-scaffold/internal/plugins"
	"code/gin-scaffold/internal/settings"
	"code/gin-scaffold/middleware"
	"code/gin-scaffold/models"
	"code/gin-scaffold/router"

	"github.com/gin-gonic/gin"
)

func init() {
	settings.SetUp()
	logging.SetUp()
	gredis.SetUp()
	database.SetUp()
	acs.Setup()
}

func main() {
	if !settings.Setting.SystemBaseConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	models.AutoMigrate()
	middleware.Init(r)
	router.Init(r)
	r.GET("/ping", plugins.Ping)
	models.AutoMigrate()
	err := r.Run(":5000")
	if err != nil {
		global.Logger.Fatal("Fail to start project", err)
	}

}
