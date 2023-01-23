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

func Infof(format string, a interface{}) {
	if consts.LogLevelInfo <= logLevel {
		fmt.Printf(format, a)
	}
}

func Error(msg ...interface{}) {
	if consts.LogLevelError <= logLevel {
		fmt.Println(msg...)
	}
}

func Errorf(format string, a interface{}) {
	if consts.LogLevelError <= logLevel {
		fmt.Printf(format, a)
	}
}

func Warn(msg ...interface{}) {
	if consts.LogLevelWarn <= logLevel {
		fmt.Println(msg...)
	}
}

func Warnf(format string, a interface{}) {
	if consts.LogLevelWarn <= logLevel {
		fmt.Printf(format, a)
	}
}

func Debug(msg ...interface{}) {
	if consts.LogLevelDebug <= logLevel {
		fmt.Println(msg...)
	}
}

func Debugf(format string, a interface{}) {
	if consts.LogLevelDebug <= logLevel {
		fmt.Printf(format, a)
	}
}

type MyLogger struct{}

func (l MyLogger) Errorf(format string, v ...interface{}) {
	Errorf(format, v)
}
func (l MyLogger) Warnf(format string, v ...interface{}) {
	Warnf(format, v)
}
func (l MyLogger) Debugf(format string, v ...interface{}) {
	Debugf(format, v)
}
