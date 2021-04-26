package repositories

import "github.com/sptGabriel/go-ddd/domain/entities"

const PersonCollection = "persons"

type PersonRepository interface {
	Save(person *entities.Person) error
	Update(person *entities.Person) error
	GetById(id string) (person *entities.Person, err error)
	GetAll() (persons []*entities.Person, err error)
	Delete(id string) error
}
