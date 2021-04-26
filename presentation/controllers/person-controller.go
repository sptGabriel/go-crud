package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sptGabriel/go-ddd/domain/commands"
	valueObjects "github.com/sptGabriel/go-ddd/domain/value-objects"
	"github.com/sptGabriel/go-ddd/infra/commandBus"
	"github.com/sptGabriel/go-ddd/utils"
)

type createPersonDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

type PersonController struct {
	commandBus commandBus.CommandBus
	name       string
}

func NewPersonController(name string, cb commandBus.CommandBus) *PersonController {
	return &PersonController{name: name, commandBus: cb}
}

func (ctr *PersonController) CreateNewPerson(ctx *fiber.Ctx) error {
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
	command := commands.CreatePersonCommand(commands.CreatePersonCommand{Name: name, Email: email, Password: password})
	value, fail := ctr.commandBus.Publish(command)
	if fail != nil {
		return utils.CtxError(ctx, fail.Code, fail.Error)
	}
	return ctx.Status(http.StatusOK).JSON(value)
}

func (*PersonController) GetPerson(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON("Hello")
}
