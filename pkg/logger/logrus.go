package logger

import "github.com/sirupsen/logrus"

func Info(msg ...interface{}) {
	logrus.Info(msg...)
}

func Error(msg ...interface{}) {
	logrus.Error(msg...)
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}
