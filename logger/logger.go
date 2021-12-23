package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error
	config := zap.NewProductionConfig()
	EncoderConfig := zap.NewProductionEncoderConfig()
	EncoderConfig.TimeKey = "timestamp"
	EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	EncoderConfig.StacktraceKey = ""
	config.EncoderConfig = EncoderConfig

	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
