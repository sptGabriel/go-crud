package handlers

import (
	"github.com/sptGabriel/go-ddd/application/errors"
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
	const op errors.Op = "person.handlers.getPerson"
	person, err := ch.repository.GetById("ad4c7d41-4e17-42c2-912b-c6a717cf69d7")
	if err != nil {
		return nil, errors.E(op, err, errors.GetCode(err))
	}
	return person, nil
}
