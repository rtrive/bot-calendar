package log

import (
	"encoding/json"

	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func init() {

	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout", "/tmp/logs"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase"
		}
	  }`)
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger := zap.Must(cfg.Build())
	defer logger.Sync()

	log = logger.Sugar()
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
