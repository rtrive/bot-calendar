package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.SugaredLogger

func init() {

	cfgEncoderConfig := zapcore.EncoderConfig{
		MessageKey:  "message",
		LevelKey:    "level",
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		TimeKey:     "timestamp",
		EncodeTime:  zapcore.RFC3339TimeEncoder,
	}
	config := zap.NewProductionConfig()
	config.EncoderConfig = cfgEncoderConfig
	tmpLog, _ := config.Build(zap.AddCallerSkip(1))
	log = tmpLog.Sugar()

}

func Info(message string) {
	log.Info(message)
}

func Error(message string) {
	log.Error(message)
}

func Debug(message string) {
	log.Debug(message)
}
