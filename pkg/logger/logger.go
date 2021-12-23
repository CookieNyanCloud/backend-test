package logger

import "github.com/sirupsen/logrus"

type Logger interface {
	Info(msg string)
	Infof(msg string, params map[string]interface{})
	Error(msg string)
	Errorf(msg string, params map[string]interface{})
}

func Info(msg ...interface{}) {
	logrus.Info(msg...)
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Error(msg ...interface{}) {
	logrus.Error(msg...)
}

func Errorf(format string, err error) {
	if err != nil {
		logrus.Errorf(format, err)
		return
	}
}
