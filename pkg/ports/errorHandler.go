package ports

type ErrorHandler interface {
	CaptureError(msg string)
}
