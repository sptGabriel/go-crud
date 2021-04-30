package person

import (
	"encoding/json"
	"fmt"

	"github.com/sptGabriel/go-ddd/application/errors"
)

const (
	minNameLength = 4
	maxNameLength = 31
)

var (
	ErrName = fmt.Errorf("Value must be between %v and %v characters long", minNameLength, maxNameLength)
)

type Name struct {
	firstName string
	lastName  string
}

func NewName(firstName string, lastName string) (Name, error) {
	var op errors.Op = "person.name"
	fullName := fmt.Sprintf("%s %s", firstName, lastName)
	if !(len(fullName) >= minNameLength && len(fullName) <= maxNameLength) || fullName == "" {
		return Name{}, errors.E(op, ErrName, errors.KindUnprocessable)
	}
	return Name{firstName, lastName}, nil
}

func (n Name) FirstName() string {
	return n.firstName
}

func (n Name) LastName() string {
	return n.lastName
}

func (n Name) FullName() string {
	return fmt.Sprintf("%s %s", n.firstName, n.lastName)
}

func (n Name) Equals(n2 Name) bool {
	return n.FullName() == n2.FullName()
}

func (n *Name) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}{
		FirstName: n.FirstName(),
		LastName:  n.LastName(),
	})
}

func (n *Name) UnmarshalText(firstName string, lastName string) error {
	var err error
	*n, err = NewName(firstName, lastName)
	return err
}
