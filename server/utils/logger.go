package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// InitLogger 初始化结构化日志记录器
func InitLogger() error {
	// 根据环境设置日志级别
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	// 设置编码器配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 设置日志级别
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zapcore.InfoLevel
	}

	// 创建核心
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		level,
	)

	// 创建日志记录器
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	// 替换全局日志记录器
	zap.ReplaceGlobals(Logger)

	return nil
}

// GetLogger 获取日志记录器实例
func GetLogger() *zap.Logger {
	if Logger == nil {
		// 如果未初始化，使用默认配置初始化
		InitLogger()
	}
	return Logger
}

// SyncLogger 刷新日志缓冲区
func SyncLogger() {
	if Logger != nil {
		Logger.Sync()
	}
}

// 便捷日志方法
func Debug(msg string, fields ...zap.Field) {
	GetLogger().Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	GetLogger().Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	GetLogger().Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	GetLogger().Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	GetLogger().Fatal(msg, fields...)
}

// 带上下文的日志方法
func DebugCtx(ctx interface{}, msg string, fields ...zap.Field) {
	GetLogger().Debug(msg, append(fields, zap.Any("context", ctx))...)
}

func InfoCtx(ctx interface{}, msg string, fields ...zap.Field) {
	GetLogger().Info(msg, append(fields, zap.Any("context", ctx))...)
}

func WarnCtx(ctx interface{}, msg string, fields ...zap.Field) {
	GetLogger().Warn(msg, append(fields, zap.Any("context", ctx))...)
}

func ErrorCtx(ctx interface{}, msg string, fields ...zap.Field) {
	GetLogger().Error(msg, append(fields, zap.Any("context", ctx))...)
}

// HTTP请求日志字段
func HTTPRequestFields(method, path, status string, duration float64) []zap.Field {
	return []zap.Field{
		zap.String("http.method", method),
		zap.String("http.path", path),
		zap.String("http.status", status),
		zap.Float64("http.duration_seconds", duration),
	}
}

// 数据库操作日志字段
func DBOperationFields(operation, collection string, duration float64) []zap.Field {
	return []zap.Field{
		zap.String("db.operation", operation),
		zap.String("db.collection", collection),
		zap.Float64("db.duration_seconds", duration),
	}
}

// 错误日志字段
func ErrorFields(err error) []zap.Field {
	return []zap.Field{
		zap.Error(err),
		zap.String("error.message", err.Error()),
	}
}
