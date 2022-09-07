package logging

import (
	"code/gin-scaffold/internal/global"
	"code/gin-scaffold/internal/settings"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)

func getLogFilename(filename string) string {
	return fmt.Sprintf(
		"%s%s_%s-%s.%s",
		settings.Setting.LoggerConfig.Path,
		settings.Setting.LoggerConfig.Prefix,
		filename,
		time.Now().Format(settings.Setting.LoggerConfig.Format),
		settings.Setting.LoggerConfig.Suffix,
	)
}

func getLogWriter(filename string) zapcore.WriteSyncer {
	// go get -u github.com/natefinch/lumberjack
	// 使用lumberjack切割日志
	lumberjackLogger := &lumberjack.Logger{
		Filename:   getLogFilename(filename),               // 日志文件的位置
		MaxSize:    settings.Setting.LoggerConfig.MaxSize,  //日志文件最大大小, 单位MB
		MaxBackups: settings.Setting.LoggerConfig.Backups,  //旧文件个数
		MaxAge:     settings.Setting.LoggerConfig.Age,      //保留最大天数
		Compress:   settings.Setting.LoggerConfig.Compress, //是否压缩、归档
	}
	return zapcore.AddSync(lumberjackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// SetUp 初始化sugar logger对象
func SetUp() {
	normalWriterSyncer := getLogWriter(settings.Setting.LoggerConfig.NormalLog)
	errorWriterSyncer := getLogWriter(settings.Setting.LoggerConfig.ErrorLog)
	encoder := getEncoder()
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, normalWriterSyncer, zapcore.InfoLevel),
		zapcore.NewCore(encoder, errorWriterSyncer, zapcore.ErrorLevel),
	)
	logger := zap.New(core, zap.AddCaller())
	global.Logger = logger.Sugar()
	log.Println("日志logger初始化完成.")
}
