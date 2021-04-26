package utils

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type JError struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func NewJError(err error) JError {
	jerr := JError{Message: "generic message", Error: "generic error"}
	if err != nil {
		jerr.Message = err.Error()
	}
	return jerr
}

func CtxError(ctx *fiber.Ctx, code int, err error) error {
	return ctx.
		Status(http.StatusUnprocessableEntity).
		JSON(NewJError(err))
}
