package handlers

import (
	"github.com/sptGabriel/go-ddd/application/errors"
	"github.com/sptGabriel/go-ddd/domain/person/commands"
	cb "github.com/sptGabriel/go-ddd/infra/commandBus"
	repo "github.com/sptGabriel/go-ddd/infra/repositories"
)

type GetPersonHandler struct {
	repository repo.PersonRepository
}

func NewGetPersonCommandHandler(r repo.PersonRepository) GetPersonHandler {
	return GetPersonHandler{repository: r}
}

func (ch GetPersonHandler) Execute(c cb.Command) (interface{}, error) {
	const op errors.Op = "person.handlers.getPerson"
	cmd, ok := c.(commands.GetPersonCommand)
	if !ok {
		return nil, errors.E(op, errors.ErrInternal, errors.KindUnexpected)
	}
	res, err := ch.repository.GetById(cmd.Id)
	if err != nil {
		return nil, errors.E(op, err, errors.GetCode(err))
	}
	return res, nil
}
