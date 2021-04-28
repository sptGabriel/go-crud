package errors

import (
	"bytes"
	"fmt"
	"net/http"
)

//type AppError struct {
//	//	Error   error
//	//	Message string
//	//	Code    int
//	//}

//	//var (
//	//	ErrInternal = errors.New("InternalError")
//	//)

type Op string

type Code int

type Message string

const (
	KindNotFound       = http.StatusNotFound
	KindUnauthorized   = http.StatusUnauthorized
	KindUnprocessable  = http.StatusUnprocessableEntity
	KindBadRequest     = http.StatusBadRequest
	KindUnexpected     = http.StatusInternalServerError
	KindEntityNotFound = http.StatusOK
)

var (
	ErrInternal = fmt.Errorf("internal error")
)

type Error struct {
	Op      Op    // operation
	Code    Code  // category of errors
	Err     error // the wrapped error
	Message Message
}

func (e *Error) Error() string {
	var buf bytes.Buffer
	if e.Op != "" {
		fmt.Fprintf(&buf, "%s: ", e.Op)
	}
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != 0 {
			fmt.Fprintf(&buf, "<%v> ", e.Code)
		}
		buf.WriteString(string(e.Error()))
	}
	return buf.String()
}

func GetCode(err error) Code {
	e, ok := err.(*Error)
	if !ok {
		return KindUnexpected
	}

	if e.Code != 0 {
		return e.Code
	}

	return GetCode(e.Err)
}

func E(args ...interface{}) error {
	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case Op:
			e.Op = arg
		case error:
			e.Err = arg
		case Code:
			e.Code = arg
		default:
			panic("bad call to E")
		}
	}

	return e
}
