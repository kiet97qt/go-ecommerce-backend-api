package loggers

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

// init configures the global zap logger instance.
func init() {
	encoder := getEncoderLog()
	writer := getWriterSync()

	core := zapcore.NewCore(encoder, writer, zapcore.InfoLevel)
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

// getEncoderLog defines how log entries are encoded (JSON format, time, level, caller, etc.).
func getEncoderLog() zapcore.Encoder {
	cfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}

	return zapcore.NewJSONEncoder(cfg)
}

// getWriterSync configures where logs are written (here: logs/app.log file).
func getWriterSync() zapcore.WriteSyncer {
	const logDir = "logs"
	const logFile = "logs/app.log"

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.MkdirAll(logDir, 0o755); err != nil {
			log.Fatalf("failed to create log directory: %v", err)
		}
	}

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}

	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(file), zapcore.AddSync(os.Stdout))
}

// GetLogger returns the shared zap logger instance for use in other packages.
func GetLogger() *zap.Logger {
	return logger
}

// Debug logs a message at DebugLevel.
func Debug(msg string, fields ...zap.Field) {
	if logger == nil {
		return
	}
	logger.Debug(msg, fields...)
}

// Info logs a message at InfoLevel.
func Info(msg string, fields ...zap.Field) {
	if logger == nil {
		return
	}
	logger.Info(msg, fields...)
}

// Warn logs a message at WarnLevel.
func Warn(msg string, fields ...zap.Field) {
	if logger == nil {
		return
	}
	logger.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel.
func Error(msg string, fields ...zap.Field) {
	if logger == nil {
		return
	}
	logger.Error(msg, fields...)
}

// Fatal logs a message at FatalLevel then calls os.Exit(1).
func Fatal(msg string, fields ...zap.Field) {
	if logger == nil {
		return
	}
	logger.Fatal(msg, fields...)
}
