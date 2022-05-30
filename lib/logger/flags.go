package logger

import (
	"flag"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logLevel = flag.String("log-level", "debug", "Log level")

	// DefaultLogger is the default logger used by application
	DefaultLogger *zap.SugaredLogger
)

// Init Initializes logger
func Init() {
	DefaultLogger = NewLogger()
}

// Flush flushes messages of logger which are written asynchronously
func Flush() {
	if DefaultLogger != nil {
		_ = DefaultLogger.Sync()
	}
}

func levelToZap() zap.AtomicLevel {
	var level zapcore.Level

	switch *logLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.DebugLevel
	}

	return zap.NewAtomicLevelAt(level)
}

// NewLogger creates a new logger instance
func NewLogger() *zap.SugaredLogger {
	config := zap.NewProductionConfig()
	config.Level = levelToZap()

	logger, _ := config.Build()
	sugar := logger.Sugar()

	return sugar
}
