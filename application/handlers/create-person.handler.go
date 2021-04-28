package handlers

import (
	"github.com/sptGabriel/go-ddd/application/errors"
	"github.com/sptGabriel/go-ddd/domain/person"
	cb "github.com/sptGabriel/go-ddd/infra/commandBus"
	repo "github.com/sptGabriel/go-ddd/infra/repositories"
)

type CreatePersonHandler struct {
	repository repo.PersonRepository
}

func NewCreatePersonCommandHandler(r repo.PersonRepository) CreatePersonHandler {
	return CreatePersonHandler{repository: r}
}

func (ch CreatePersonHandler) Execute(c cb.Command) (*person.Person, error) {
	const op = "person.handlers.getPerson"
	person, err := ch.repository.GetById("111111")
	errors.E(op, err)
	return person, nil
}
