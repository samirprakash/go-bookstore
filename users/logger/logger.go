package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	c := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseColorLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	if Log, err := c.Build(); err != nil {
		panic(err)
	} else {
		zap.ReplaceGlobals(Log)
	}
}

func GetLogger() *zap.Logger {
	return log
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
	log.Sync()
}

func Error(msg string, err error, fields ...zap.Field) {
	fields = append(fields, zap.NamedError("error", err))
	log.Error(msg, fields...)
	log.Sync()
}
