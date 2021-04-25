package valueObjects

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var (
	minPasswordLength  = 8
	ErrPassword        = errors.New("password: does not meet the minimum criteria")
	ErrInvalidPassword = errors.New("password: does not match")
)

type Password struct {
	value string
}

func NewPassword(p string) (Password, error) {
	if len(p) < minPasswordLength || p == "" {
		return Password{}, fmt.Errorf("invalid password, min 8 characters and non-empty")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return Password{}, errors.New("error when generating the hash")
	}
	return Password{value: string(hash)}, nil
}

func (p Password) Value() string {
	return p.value
}

func (p Password) Equals(p2 Password) bool {
	return p.value == p2.value
}

func (p Password) ComparePassword(v string) bool {
	plainPwd := []byte(v)
	entityPwd := []byte(p.value)
	err := bcrypt.CompareHashAndPassword(entityPwd, plainPwd)
	return err == nil
}
