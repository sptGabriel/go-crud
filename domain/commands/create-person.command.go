package commands

import person "github.com/sptGabriel/go-ddd/domain/entities"

type CreatePersonCommand struct {
	Name     person.Name
	Email    person.Email
	Password person.Password
}
