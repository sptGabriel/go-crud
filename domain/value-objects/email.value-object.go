package valueObjects

import (
	"errors"
	"net/mail"
	"strings"
)

var (
	ErrEmail = errors.New("email: could not use invalid email")
)

type email struct {
	Address string `json: "email"`
}

type Email struct {
	email
}

func NewEmail(address string) (Email, error) {
	isValid := valid(address)
	if !isValid {
		return Email{}, ErrEmail
	}
	email := email{Address: address}
	return Email{email}, nil
}

func (e Email) Value() string {
	return e.email.Address
}

func (e Email) Format() string {
	return strings.TrimSpace(strings.ToLower(e.email.Address))
}

func (e Email) Equals(e2 Email) bool {
	return e.Value() == e2.Value()
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
