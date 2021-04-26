package appErrors

import "errors"

type AppError struct {
	Error   error
	Message string
	Code    int
}

var (
	ErrInternal = errors.New("InternalError")
)
