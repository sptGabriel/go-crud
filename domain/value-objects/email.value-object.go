package valueObjects

import (
	"errors"
	"net/mail"
	"strings"
)

var (
	ErrEmail = errors.New("email: could not use invalid email")
)

type Email struct {
	address string
}

func NewEmail(address string) (Email, error) {
	isValid := valid(address)
	if !isValid {
		return Email{}, ErrEmail
	}
	return Email{address}, nil
}

func (e Email) Value() string {
	return e.address
}

func (e Email) Format() string {
	return strings.TrimSpace(strings.ToLower(e.Value()))
}

func (e Email) Equals(e2 Email) bool {
	return e.Value() == e2.Value()
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
