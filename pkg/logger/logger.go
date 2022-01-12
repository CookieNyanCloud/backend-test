package logger

import "github.com/sirupsen/logrus"

//interface for logs
type Logger interface {
	Info(msg ...interface{})
	Error(err error)
	Errorf(format string, err error)
}

//info string
func Info(msg ...interface{}) {
	logrus.Info(msg)
}

//error string
func Error(err error) {
	logrus.Error(err)
}

//error with params
func Errorf(format string, err error) {
	if err != nil {
		logrus.Errorf(format, err)
		return
	}
}
