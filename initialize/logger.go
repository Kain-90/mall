package initialize

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"mall/global"
)

func SetupLogger() *zap.Logger {
	var logger *zap.Logger

	if global.GVA_CONFIG.App.Env == "dev" {
		cfg := zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		logger, _ = cfg.Build()
	} else {
		logger, _ = zap.NewProduction()
	}
	return logger
}
