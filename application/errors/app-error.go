package errors

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

type Op string

type Code int

type Message string

const (
	KindNotFound       Code = http.StatusNotFound
	KindUnauthorized   Code = http.StatusUnauthorized
	KindUnprocessable  Code = http.StatusUnprocessableEntity
	KindBadRequest     Code = http.StatusBadRequest
	KindUnexpected     Code = http.StatusInternalServerError
	KindEntityNotFound Code = http.StatusOK
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

func GetErrorMessage(err error) string {
	e, ok := err.(*Error)
	if !ok {
		return "Internal Error"
	}
	msgarray := strings.Split(e.Error(), ":")
	msg := msgarray[len(msgarray)-1]
	return strings.Join(strings.Fields(strings.TrimSpace(msg)), " ")
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
