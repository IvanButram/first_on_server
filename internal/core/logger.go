package core

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field struct {
	Key   string
	Value any
}

type Logger struct {
	zap *zap.Logger
}

func NewLogger(logToFile bool, filePath string) (*Logger, error) {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.LevelKey = "level"
	encoderConfig.MessageKey = "msg"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder

	var writeSyncer zapcore.WriteSyncer

	if logToFile {
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return nil, fmt.Errorf("open log file: %w", err)
		}
		writeSyncer = zapcore.AddSync(file)
	} else {
		writeSyncer = zapcore.AddSync(os.Stdout)
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		writeSyncer,
		zapcore.DebugLevel,
	)

	z := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return &Logger{zap: z}, nil
}

func (l *Logger) Log(level string, message string, fields ...Field) {
	zapFields := make([]zap.Field, 0, len(fields))

	for _, field := range fields {
		zapFields = append(zapFields, zap.Any(field.Key, field.Value))
	}

	switch strings.ToLower(level) {
	case "debug":
		l.zap.Debug(message, zapFields...)
	case "info":
		l.zap.Info(message, zapFields...)
	case "warn":
		l.zap.Warn(message, zapFields...)
	case "error":
		l.zap.Error(message, zapFields...)
	default:
		l.zap.Info(message, append(zapFields, zap.String("unknown_level", level))...)
	}
}

func (l *Logger) Debug(message string, fields ...Field) {
	l.Log("debug", message, fields...)
}

func (l *Logger) Info(message string, fields ...Field) {
	l.Log("info", message, fields...)
}

func (l *Logger) Warn(message string, fields ...Field) {
	l.Log("warn", message, fields...)
}

func (l *Logger) Error(message string, fields ...Field) {
	l.Log("error", message, fields...)
}

func (l *Logger) Sync() error {
	return l.zap.Sync()
}
