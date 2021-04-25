package utils

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CtxError(ctx *fiber.Ctx, code int, err error) error {
	return ctx.
		Status(http.StatusUnprocessableEntity).
		JSON(err.Error())
}
