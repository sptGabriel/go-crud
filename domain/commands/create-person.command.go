package commands

import valueObjects "github.com/sptGabriel/go-ddd/domain/value-objects"

type CreatePersonCommand struct {
	Name     valueObjects.Name
	Email    valueObjects.Email
	Password valueObjects.Password
}
