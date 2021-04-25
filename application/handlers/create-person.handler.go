package handlers

import (
	"errors"
	"fmt"

	"github.com/sptGabriel/go-ddd/domain/commands"
	"github.com/sptGabriel/go-ddd/infra/commandBus"
)

type CreatePersonHandler struct {
}

func NewCreatePersonCommandHandler() CreatePersonHandler {
	return CreatePersonHandler{}
}

func (ch CreatePersonHandler) Handle(c commandBus.Command) (commands.CreatePerson, error) {
	cmd, ok := c.(commands.CreatePerson)
	if !ok {
		return nil, errors.New("Invalid command")
	}
	fmt.Println(cmd)
	return commands.CreatePerson{FirstName: "a", LastName: "b"}, nil
}
