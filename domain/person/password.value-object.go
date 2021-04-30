package person

import (
	"encoding/json"
	"fmt"

	"github.com/sptGabriel/go-ddd/application/errors"
	"golang.org/x/crypto/bcrypt"
)

const (
	minPasswordLength = 8
)

var (
	ErrPassword        = fmt.Errorf("does not meet the minimum criteria")
	ErrPasswordMin     = fmt.Errorf("Value must be at least %v characters long", minPasswordLength)
	ErrInvalidPassword = fmt.Errorf("password does not match")
	ErrPasswordHash    = fmt.Errorf("error when generating the hash")
)

type Password struct {
	value string
}

func NewPassword(p string) (Password, error) {
	var op errors.Op = "person.password"
	if len(p) < minPasswordLength {
		return Password{}, errors.E(op, ErrPasswordMin, errors.KindUnprocessable)
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return Password{}, errors.E(op, ErrPasswordHash, errors.KindUnexpected)
	}
	return Password{value: string(hash)}, nil
}

func (p Password) Value() string {
	return p.value
}

func (p Password) Equals(p2 Password) bool {
	return p.value == p2.value
}

func (p *Password) UnmarshalText(pwd string) error {
	var err error
	*p, err = NewPassword(pwd)
	return err
}

func (e *Password) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Value())
}

func (p Password) ComparePassword(v string) bool {
	plainPwd := []byte(v)
	entityPwd := []byte(p.value)
	err := bcrypt.CompareHashAndPassword(entityPwd, plainPwd)
	return err == nil
}
