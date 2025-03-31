package autoload

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func InitZapLogger() *zap.Logger {
	// 日志轮转配置
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    100,  // MB
		MaxBackups: 30,   // 保留旧文件数量
		MaxAge:     28,   // 保留天数
		Compress:   true, // 压缩旧日志
	}
	// Zap 配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 时间格式
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // JSON 格式
		zapcore.AddSync(lumberjackLogger),     // 输出到文件
		zap.NewAtomicLevelAt(zap.InfoLevel),   // 日志级别
	)
	// 添加控制台输出（可选）
	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		zap.NewAtomicLevelAt(zap.DebugLevel),
	)
	// 组合多个输出
	teeCore := zapcore.NewTee(core, consoleCore)
	return zap.New(teeCore, zap.AddCaller()) // 添加调用方信息
}
