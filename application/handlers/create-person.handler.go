package handlers

import (
	appErrors "github.com/sptGabriel/go-ddd/application/errors"
	"github.com/sptGabriel/go-ddd/domain/commands"
	"github.com/sptGabriel/go-ddd/domain/entities"
	"github.com/sptGabriel/go-ddd/domain/repositories"
	"github.com/sptGabriel/go-ddd/infra/commandBus"
)

type CreatePersonHandler struct {
	repository repositories.PersonRepository
}

func NewCreatePersonCommandHandler(r repositories.PersonRepository) CreatePersonHandler {
	return CreatePersonHandler{repository: r}
}

func (ch CreatePersonHandler) Handle(c commandBus.Command) (interface{}, *commandBus.CommandHandlerError) {
	cmd, ok := c.(commands.CreatePersonCommand)
	if !ok {
		return nil, &commandBus.CommandHandlerError{Error: appErrors.ErrInternal, Code: 500}
	}
	person := entities.NewPerson(cmd.Name, cmd.Email, cmd.Password)
	if err := ch.repository.Save(person); err != nil {
		return nil, &commandBus.CommandHandlerError{Error: appErrors.ErrInternal, Code: 500}
	}
	return person.ToJSON(), &commandBus.CommandHandlerError{}
}
