package entities

import (
	"encoding/json"
	"fmt"
	"time"

	valueObject "github.com/sptGabriel/go-ddd/domain/value-objects"
)

type Person struct {
	Id       string
	Name     valueObject.Name
	Email    valueObject.Email
	Password valueObject.Password
	Created  time.Time
	Updated  time.Time
}

func NewPerson(name valueObject.Name, email valueObject.Email, password valueObject.Password) *Person {
	return &Person{
		Id:       "a",
		Name:     name,
		Email:    email,
		Password: password,
		Created:  time.Now(),
	}
}

func (p *Person) ToJSON() {
	personJSON, err := json.Marshal(p)
	fmt.Println(personJSON, err)
}
