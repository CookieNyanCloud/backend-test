package logger

import "github.com/sirupsen/logrus"

//interface for logs
type Logger interface {
	Info(msg string)
	Infof(msg string, params map[string]interface{})
	Error(msg string)
	Errorf(msg string, params map[string]interface{})
}

//info string
func Info(msg ...interface{}) {
	logrus.Info(msg...)
}

//info with params
func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

//error string
func Error(msg ...interface{}) {
	logrus.Error(msg...)
}

//error with params
func Errorf(format string, err error) {
	if err != nil {
		logrus.Errorf(format, err)
		return
	}
}
