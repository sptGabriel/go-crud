package valueObjects

import (
	"errors"
	"fmt"
)

const (
	minNameLength = 4
	maxNameLength = 50
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
	fullName := firstName + lastName
	if !(len(fullName) >= minNameLength && len(fullName) <= maxNameLength) || fullName == "" {
		return Name{}, fmt.Errorf("invalid name, must be with 20 charcters and non-empty")
	}
	return Name{firstName, lastName}, nil
}

func (n Name) FullName() string {
	return string(n.firstName + n.lastName)
}

func (n Name) Equals(n2 Name) bool {
	return n.FullName() == n2.FullName()
}
