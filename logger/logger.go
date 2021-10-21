package logger

import (
	"context"
	contextx "gin-essential/ctx"
	"os"

	"github.com/google/wire"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Define key
const (
	TraceIDKey = "trace_id"
	UserIDKey  = "user_id"
)

// LoggerSet 日志
var LoggerSet = wire.NewSet(InitLogger)

// Logger 实例
var Logger *zap.Logger

// InitLogger 初始化日志
func InitLogger() *zap.Logger {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	Logger = zap.New(core)
	Debug = Logger.Debug
	Info = Logger.Info
	Warn = Logger.Warn
	Error = Logger.Error
	Fatal = Logger.Fatal
	Panic = Logger.Panic
	DPanic = Logger.DPanic

	return Logger

}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

// WithContext Use context create entry
func WithContext(ctx context.Context) map[string]interface{} {
	if ctx == nil {
		ctx = context.Background()
	}

	fields := map[string]interface{}{}
	if v, ok := contextx.FromTraceID(ctx); ok {
		fields[TraceIDKey] = v
	}

	if v, ok := contextx.FromUserID(ctx); ok {
		fields[UserIDKey] = v
	}

	return fields
}

// Define alias
var (
	Debug  = Logger.Debug
	Info   = Logger.Info
	Warn   = Logger.Warn
	Error  = Logger.Error
	Fatal  = Logger.Fatal
	Panic  = Logger.Panic
	DPanic = Logger.DPanic
)
