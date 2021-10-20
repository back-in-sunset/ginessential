package logger

import (
	"github.com/google/wire"
	"go.uber.org/zap"
)

// LoggerSet 日志
var LoggerSet = wire.NewSet(InitLogger)

// Logger 实例
var Logger *zap.Logger

// InitLogger 初始化日志
func InitLogger() *zap.Logger {
	Logger, _ = zap.NewDevelopment()
	return Logger
}
