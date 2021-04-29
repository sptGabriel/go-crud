package commands

import "github.com/sptGabriel/go-ddd/domain/person"

type CreatePersonCommand struct {
	Name     person.Name
	Email    person.Email
	Password person.Password
}
