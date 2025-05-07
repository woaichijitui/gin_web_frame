package core

import (
	"gin_web_frame/core/internal"
	"gin_web_frame/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
)

func ZapInit() (logger *zap.Logger) {
	// 设置日志前缀
	internal.SetLogPrefix()

	cfg := zap.NewDevelopmentConfig()
	// 设置日志等级
	//cfg.Level = zap.NewAtomicLevelAt(global.CONFIG.Zap.SetLevel())
	// 设置日志时间格式化
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	// 设置颜色
	//cfg.EncoderConfig.EncodeLevel = coloredLevelEncoder
	// 设置颜色
	cfg.EncoderConfig.EncodeLevel = global.CONFIG.Zap.LevelEncoder()
	//设置栈命
	cfg.EncoderConfig.StacktraceKey = global.CONFIG.Zap.StacktraceKey

	// 创建自定义的 Encoder
	encoder := &internal.MyEncoder{
		LogRootDir: filepath.Join("./", global.CONFIG.Zap.Director),
		Encoder:    zapcore.NewConsoleEncoder(cfg.EncoderConfig), // 使用 Console 编码器
	}

	// 自定义日志输出 日志等级
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), global.CONFIG.Zap.SetLevel())

	// 是否显示行号
	// 默认不显示
	if global.CONFIG.Zap.ShowLine {
		logger = zap.New(core, zap.AddCaller())
	} else {
		logger = zap.New(core)
	}

	return logger
}
