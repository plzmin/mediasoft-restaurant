package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Interface interface {
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
	Fatal(message string, args ...interface{})
}

type Logger struct {
	logger *logrus.Entry
}

func New() *Logger {
	logger := logrus.NewEntry(logrus.StandardLogger())
	logger.Logger.SetFormatter(&logrus.JSONFormatter{})
	return &Logger{logger: logger}
}

func (l *Logger) Info(message string, args ...interface{}) {
	l.logger.Infof(message, args...)
}

func (l *Logger) Debug(message string, args ...interface{}) {
	l.logger.Debugf(message, args...)
}

func (l *Logger) Warn(message string, args ...interface{}) {
	l.logger.Warnf(message, args...)
}

func (l *Logger) Error(message string, args ...interface{}) {
	l.logger.Errorf(message, args...)
}

func (l *Logger) Fatal(message string, args ...interface{}) {
	l.logger.Fatalf(message, args...)
	os.Exit(1)
}
