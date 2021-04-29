package utils

import (
	"fmt"

	"github.com/sptGabriel/go-ddd/application/errors"
)

type JError struct {
	Message string `json:"message"`
}

func NewJError(err error) JError {
	jerr := JError{Message: "Internal Error"}
	e, ok := err.(*errors.Error)
	if !ok {
		fmt.Println(ok, "ok")
		return jerr
	}
	if e.Err != nil {
		jerr.Message = errors.GetErrorMessage(err)
	}
	return jerr
}
