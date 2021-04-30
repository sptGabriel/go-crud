package person

import (
	"encoding/json"
	"fmt"
	"net/mail"
	"strings"

	"github.com/sptGabriel/go-ddd/application/errors"
)

var (
	ErrEmail = fmt.Errorf("email: could not use invalid email")
)

type Email struct {
	value string
}

func NewEmail(address string) (Email, error) {
	var op errors.Op = "person.email"
	isValid := valid(address)
	if !isValid {
		return Email{}, errors.E(op, ErrEmail, errors.KindUnprocessable)
	}
	return Email{address}, nil
}

func (e Email) Value() string {
	return e.value
}

func (e Email) Format() string {
	return strings.TrimSpace(strings.ToLower(e.Value()))
}

func (e Email) Equals(e2 Email) bool {
	return e.Value() == e2.Value()
}

func (e *Email) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Value())
}

func (n *Email) UnmarshalText(d string) error {
	var err error
	*n, err = NewEmail(d)
	return err
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
