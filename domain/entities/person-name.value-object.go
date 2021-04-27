package person

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	minNameLength = 4
	maxNameLength = 31
)

var (
	ErrName     = errors.New("name: could not use invalid Name")
	ErrTooShort = fmt.Errorf("%w: min length allowed is %d", ErrName, minNameLength)
	ErrTooLong  = fmt.Errorf("%w: max length allowed is %d", ErrName, maxNameLength)
)

type Name struct {
	firstName string
	lastName  string
}

func NewName(firstName string, lastName string) (Name, error) {
	fullName := fmt.Sprintf("%s %s", firstName, lastName)
	if !(len(fullName) >= minNameLength && len(fullName) <= maxNameLength) || fullName == "" {
		return Name{}, fmt.Errorf("invalid name, must be with 20 charcters and non-empty")
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
