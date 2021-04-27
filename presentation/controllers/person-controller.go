package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	cmd "github.com/sptGabriel/go-ddd/domain/commands"
	person "github.com/sptGabriel/go-ddd/domain/entities"
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
	email, err := person.NewEmail(dto.Email)
	if err != nil {
		return utils.CtxError(ctx, http.StatusUnprocessableEntity, err)
	}
	name, err := person.NewName(dto.FirstName, dto.LastName)
	if err != nil {
		return utils.CtxError(ctx, http.StatusUnprocessableEntity, err)
	}
	password, err := person.NewPassword(dto.Password)
	if err != nil {
		return utils.CtxError(ctx, http.StatusUnprocessableEntity, err)
	}
	command := cmd.CreatePersonCommand(cmd.CreatePersonCommand{Name: name, Email: email, Password: password})
	value, fail := ctr.commandBus.Publish(command)
	if value == nil {
		return utils.CtxError(ctx, fail.Code, fail.Error)
	}
	return ctx.Status(http.StatusOK).JSON(value)
}

func (*PersonController) GetPerson(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON("Hello")
}
