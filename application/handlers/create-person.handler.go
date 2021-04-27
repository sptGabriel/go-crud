package handlers

import (
	"github.com/google/uuid"
	"github.com/sptGabriel/go-ddd/domain/commands"
	vo "github.com/sptGabriel/go-ddd/domain/entities"
	cb "github.com/sptGabriel/go-ddd/infra/commandBus"
	repo "github.com/sptGabriel/go-ddd/infra/repositories"
	"github.com/sptGabriel/go-ddd/utils"
)

type CreatePersonHandler struct {
	repository repo.PersonRepository
}

func NewCreatePersonCommandHandler(r repo.PersonRepository) CreatePersonHandler {
	return CreatePersonHandler{repository: r}
}

func (ch CreatePersonHandler) Execute(c cb.Command) *utils.Either {
	//items, err := ch.repository.GetAll()
	//if err != nil {
	//	return nil, &appErrors.AppError{Error: appErrors.ErrInternal, Code: 500}
	//}
	cmd, ok := c.(commands.CreatePersonCommand)
	if !ok {
		return &utils.Either{}
	}
	person := vo.NewPerson(uuid.New().String(), cmd.Name, cmd.Email, cmd.Password)
	if err := ch.repository.Save(person); err != nil {
		return &utils.Either{}
	}
	return &utils.Either{}
}
