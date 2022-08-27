package native

import (
	logger "log"

	"git.snapp.ninja/search-and-discovery/framework/pkg/ports"
)

type NativeLogger struct {
}

func New() ports.Logger {
	return &NativeLogger{}
}

func (nl *NativeLogger) Info(msg string, params ...interface{}) {
	logger.Printf(msg, params...)
}

func (nl *NativeLogger) Warn(msg string, params ...interface{}) {
	logger.Printf(msg, params...)
}

func (nl *NativeLogger) Error(msg string, params ...interface{}) {
	logger.Printf(msg, params...)
}
