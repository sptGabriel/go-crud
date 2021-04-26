package repositoriesImpl

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sptGabriel/go-ddd/domain/entities"
	"github.com/sptGabriel/go-ddd/domain/repositories"
)

type personRepository struct {
	tableName string
	conn      *pgxpool.Pool
}

func NewPersonRepository(conn *pgxpool.Pool) repositories.PersonRepository {
	return &personRepository{tableName: repositories.PersonCollection, conn: conn}
}

func (r *personRepository) Save(p *entities.Person) error {
	qry := `insert into persons (id, first_name, last_name, email, password, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.conn.Exec(context.Background(), qry, p.Id, p.Name.FirstName(), p.Name.LastName(), p.Email.Value(), p.Password.Value(), p.Created, p.Updated)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (r *personRepository) Update(person *entities.Person) error {
	return fmt.Errorf("err")
}

func (r *personRepository) GetById(id string) (person *entities.Person, err error) {
	return nil, fmt.Errorf("err")
}

func (r *personRepository) GetAll() (persons []*entities.Person, err error) {
	//qry := `select id, first_name, last_name, email, password created_at, updated_at from persons`
	//rows, err := r.conn.Query(context.Background(), qry)

	return nil, fmt.Errorf("err")
}

func (r *personRepository) Delete(id string) error {
	return fmt.Errorf("err")
}
