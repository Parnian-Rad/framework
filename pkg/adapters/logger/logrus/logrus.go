package logrus

import (
	"git.snapp.ninja/search-and-discovery/framework/pkg/ports"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
}

func New() ports.Logger {
	return &LogrusLogger{}
}

func (nl *LogrusLogger) Info(msg string, params ...interface{}) {
	logrus.Infof(msg, params...)
}

func (nl *LogrusLogger) Warn(msg string, params ...interface{}) {
	logrus.Warnf(msg, params...)
}

func (nl *LogrusLogger) Error(msg string, params ...interface{}) {
	logrus.Errorf(msg, params...)
}
