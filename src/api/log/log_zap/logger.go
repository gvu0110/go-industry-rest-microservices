package log_zap

import (
	"fmt"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

var (
	Log *zap.Logger
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error
	Log, err = logConfig.Build()
	if err != nil {
		panic(err)
	}
}

// Info function
func Info(message string, tags ...zap.Field) {
	Log.Info(message, tags...)
	Log.Sync()
}

// Error function
func Error(message string, err error, tags ...zap.Field) {
	message = fmt.Sprintf("%s - ERROR - %v", message, err)
	Log.Error(message, tags...)
	Log.Sync()
}

// Debug function
func Debug(message string, tags ...zap.Field) {
	Log.Debug(message, tags...)
	Log.Sync()
}

// Field function to access zap.Any without importing
func Field(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}
