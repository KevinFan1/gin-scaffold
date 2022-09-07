package database

import (
	"code/gin-scaffold/internal/global"
	"code/gin-scaffold/internal/settings"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetUp() {
	var err error

	sqlLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Millisecond * 200,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	tmpConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: settings.Setting.MysqlConfig.DisableConstraint,
		Logger:                                   sqlLogger,
	}
	global.DB, err = gorm.Open(mysql.Open(settings.Setting.MysqlConfig.DSN()), tmpConfig)

	//是否开启log显示
	if settings.Setting.MysqlConfig.ShowLog {
		global.DB = global.DB.Debug()
	}

	if err != nil {
		log.Panicf("开启连接mysql数据库失败: %s", err)
	}

	sqlDB, err := global.DB.DB()
	if err != nil {
		log.Panicf("开启连接mysql数据库连接池失败: %s", err)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = sqlDB.Ping()
	if err != nil {
		log.Panicf("连接mysql数据库失败. Ping: %s", err)
	}

	log.Printf("连接Mysql成功.DSN: %s", settings.Setting.MysqlConfig.DSN())
}
