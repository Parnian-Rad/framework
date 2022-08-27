package sentry

import (
	"git.snapp.ninja/search-and-discovery/framework/pkg/ports"
	"github.com/getsentry/sentry-go"
)

type sentryErrorHandler struct {
	address string
}

func New(address string) ports.ErrorHandler {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              address,
		Environment:      "",
		Debug:            false,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		panic(err.Error())
	}

	return &sentryErrorHandler{}
}

func (se *sentryErrorHandler) CaptureError(msg string) {
	sentry.CaptureMessage(msg)
}
