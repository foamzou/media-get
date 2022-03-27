package logger

import (
	"fmt"

	"github.com/foamzou/audio-get/consts"
)

var logLevel = consts.LogLevelInfo

func SetLogLevel(level string) {
	switch level {
	case "silence":
		logLevel = consts.LogLevelSilence
	case "error":
		logLevel = consts.LogLevelError
	case "warn":
		logLevel = consts.LogLevelWarn
	case "info":
		logLevel = consts.LogLevelInfo
	case "debug":
		logLevel = consts.LogLevelDebug
	default:
		panic("invalid logger level")
	}
}

func Info(msg ...interface{}) {
	if consts.LogLevelInfo <= logLevel {
		fmt.Println(msg...)
	}
}

func Error(msg ...interface{}) {
	if consts.LogLevelError <= logLevel {
		fmt.Println(msg...)
	}
}

func Warn(msg ...interface{}) {
	if consts.LogLevelWarn <= logLevel {
		fmt.Println(msg...)
	}
}

func Debug(msg ...interface{}) {
	if consts.LogLevelDebug <= logLevel {
		fmt.Println(msg...)
	}
}
