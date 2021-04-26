package entities

import (
	"time"

	"github.com/google/uuid"
	valueObject "github.com/sptGabriel/go-ddd/domain/value-objects"
	"github.com/sptGabriel/go-ddd/utils"
)

type personJson struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type Person struct {
	Id       string
	Name     valueObject.Name
	Email    valueObject.Email
	Password valueObject.Password
	Created  time.Time
	Updated  time.Time
}

func NewPerson(name valueObject.Name, email valueObject.Email, password valueObject.Password) *Person {
	id := uuid.New()
	return &Person{
		Id:       id.String(),
		Name:     name,
		Email:    email,
		Password: password,
		Created:  time.Now(),
		Updated:  time.Now(),
	}
}

func (p Person) ToJSON() personJson {
	email := p.Email.Value()
	name := p.Name.FullName()
	createdAt := p.Created.Format(utils.DateLayout)
	updatedAt := p.Updated.Format(utils.DateLayout)
	person := personJson{Id: p.Id, Email: email, Name: name, CreatedAt: createdAt, UpdatedAt: updatedAt}
	return person
}
