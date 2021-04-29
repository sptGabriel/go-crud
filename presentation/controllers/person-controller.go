package controllers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sptGabriel/go-ddd/application/errors"
	"github.com/sptGabriel/go-ddd/domain/person"
	cmd "github.com/sptGabriel/go-ddd/domain/person/commands"
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
	}
	email, err := person.NewEmail(dto.Email)
	if err != nil {
	}
	name, err := person.NewName(dto.FirstName, dto.LastName)
	if err != nil {
	}
	password, err := person.NewPassword(dto.Password)
	if err != nil {
	}
	command := cmd.CreatePersonCommand(cmd.CreatePersonCommand{Name: name, Email: email, Password: password})
	person, err := ctr.commandBus.Publish(command)
	if err != nil {
		log.Println(err)
		return ctx.Status(int(errors.GetCode(err))).JSON(utils.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(person)
}

func (*PersonController) GetPerson(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON("Hello")
}
