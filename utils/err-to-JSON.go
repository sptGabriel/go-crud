package utils

import (
	"github.com/sptGabriel/go-ddd/application/errors"
)

type JError struct {
	Message string `json:"message"`
}

func NewJError(err error) JError {
	jerr := JError{Message: "Internal Error"}
	e, _ := err.(*errors.Error)
	if e.Err != nil {
		jerr.Message = errors.GetErrorMessage(err)
	}
	return jerr
}
