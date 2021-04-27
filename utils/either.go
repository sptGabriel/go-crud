package utils

import appErrors "github.com/sptGabriel/go-ddd/application/errors"

type Either struct {
	sucess  interface{}
	failure appErrors.AppError
}
