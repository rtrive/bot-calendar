package log

import (
	u "github.com/rtrive/bot-calendar/utility"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.SugaredLogger

func logLevel() zapcore.Level {
	level := u.CheckEnv("LOG_LEVEL")
	switch level {
	case "DEBUG":
		return zap.DebugLevel
	case "INFO":
		return zap.InfoLevel
	case "ERROR":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func init() {

	cfgEncoderConfig := zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "level",
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		TimeKey:      "timestamp",
		EncodeTime:   zapcore.RFC3339TimeEncoder,
		CallerKey:    "caller",
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
	config := zap.NewProductionConfig()
	config.Level.SetLevel(logLevel())
	config.EncoderConfig = cfgEncoderConfig
	tmpLog, _ := config.Build(zap.AddCallerSkip(1))
	log = tmpLog.Sugar()

}

func Info(message string) {
	log.Info(message)
}

func Error(message error) {
	log.Error(message)
}

func Debug(message string) {
	log.Debug(message)
}
