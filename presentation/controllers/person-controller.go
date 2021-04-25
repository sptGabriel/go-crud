package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sptGabriel/go-ddd/domain/entities"
	valueObjects "github.com/sptGabriel/go-ddd/domain/value-objects"
	"github.com/sptGabriel/go-ddd/utils"
)

type createPersonDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

type PersonController struct {
	name string
}

func NewPersonController(name string) *PersonController {
	return &PersonController{name}
}

func (*PersonController) CreateNewPerson(ctx *fiber.Ctx) error {
	var dto createPersonDTO
	if err := ctx.BodyParser(&dto); err != nil {
		return utils.CtxError(ctx, http.StatusUnprocessableEntity, err)
	}
	email, err := valueObjects.NewEmail(dto.Email)
	if err != nil {
		return utils.CtxError(ctx, http.StatusUnprocessableEntity, err)
	}
	name, err := valueObjects.NewName(dto.FirstName, dto.LastName)
	if err != nil {
		return utils.CtxError(ctx, http.StatusUnprocessableEntity, err)
	}
	password, err := valueObjects.NewPassword(dto.Password)
	if err != nil {
		return utils.CtxError(ctx, http.StatusUnprocessableEntity, err)
	}
	person := entities.NewPerson(name, email, password)
	return ctx.Status(http.StatusOK).JSON("a")
}

func (*PersonController) GetPerson(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON("Hello")
}
