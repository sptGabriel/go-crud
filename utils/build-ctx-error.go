package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sptGabriel/go-ddd/application/errors"
)

func CtxError(ctx *fiber.Ctx, err error) error {
	return ctx.
		Status(int(errors.GetCode(err))).JSON(NewJError(err))
}
