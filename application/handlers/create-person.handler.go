package handlers

import (
	"github.com/google/uuid"
	"github.com/sptGabriel/go-ddd/application/errors"
	"github.com/sptGabriel/go-ddd/domain/person"
	"github.com/sptGabriel/go-ddd/domain/person/commands"
	cb "github.com/sptGabriel/go-ddd/infra/commandBus"
	repo "github.com/sptGabriel/go-ddd/infra/repositories"
)

type CreatePersonHandler struct {
	repository repo.PersonRepository
}

func NewCreatePersonCommandHandler(r repo.PersonRepository) CreatePersonHandler {
	return CreatePersonHandler{repository: r}
}

func (ch CreatePersonHandler) Execute(c cb.Command) (interface{}, error) {
	const op errors.Op = "person.handlers.createPerson"
	cmd, ok := c.(commands.CreatePersonCommand)
	if !ok {
		return nil, errors.E(op, errors.ErrInternal, errors.KindUnexpected)
	}
	hasPerson, err := ch.repository.GetByEmail(cmd.Email.Value())
	if hasPerson != nil || err != nil {
		return nil, errors.E(op, err, errors.GetCode(err))
	}
	person := person.NewPerson(uuid.New().String(), cmd.Name, cmd.Email, cmd.Password)
	if err := ch.repository.Save(person); err != nil {
		return nil, errors.E(op, err, errors.GetCode(err))
	}
	return person, nil
}
