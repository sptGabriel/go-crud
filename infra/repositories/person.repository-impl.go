package repositories

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sptGabriel/go-ddd/application/errors"
	"github.com/sptGabriel/go-ddd/domain/person"
)

type PersonRepository struct {
	conn *pgxpool.Pool
}

var (
	ErrPersonNotFound = fmt.Errorf("The person was not found")
)

func NewPersonRepository(conn *pgxpool.Pool) PersonRepository {
	return PersonRepository{conn: conn}
}

func (r *PersonRepository) Save(p *person.Person) error {
	qry := `insert into persons (id, first_name, last_name, email, password, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.conn.Exec(context.Background(), qry, p.Id, p.Name.FirstName(), p.Name.LastName(), p.Email.Value(), p.Password.Value(), time.Now(), time.Now())
	if err != nil {
		log.Println(err)
	}
	return err
}

func (r *PersonRepository) Update(person *person.Person) error {
	return fmt.Errorf("err")
}

func (r *PersonRepository) GetById(id string) (p *person.Person, err error) {
	const op errors.Op = "person.repository.getById"
	qry := `select first_name, last_name, email, password from persons WHERE id = $1`
	var person *person.Person
	var firstName, lastName, mail, pwd string
	err = r.conn.QueryRow(context.Background(), qry, id).
		Scan(&firstName, &lastName, &mail, &pwd)
	if err != nil && err != pgx.ErrNoRows {
		return nil, errors.E(op, errors.ErrInternal)
	}
	notFound := err == pgx.ErrNoRows || os.IsNotExist(err)
	if notFound {
		return nil, nil
	}
	person, err = r.unmarshalPerson(id, firstName, lastName, mail, pwd)
	if err != nil {
		return nil, errors.E(op, errors.ErrInternal)
	}
	return person, nil
}

func (r *PersonRepository) GetAll() (result []*person.Person, err error) {
	qry := `select id, first_name, last_name, email, password from persons`
	rows, err := r.conn.Query(context.Background(), qry)
	var persons []*person.Person
	if err != nil {
		if err == pgx.ErrNoRows {
			return persons, nil
		}
		log.Printf("can't get list person: %v\n", err)
		return persons, err
	}
	for rows.Next() {
		var id, firstName, lastName, mail, pwd string
		err = rows.Scan(&id, &firstName, &lastName, &mail, &pwd)
		if err != nil {
			log.Printf("Failed to build item: %v\n", err)
			return persons, err
		}
		pr, err := r.unmarshalPerson(id, firstName, lastName, mail, pwd)
		if err != nil {
			return nil, err
		}
		persons = append(persons, pr)
	}
	defer rows.Close()
	return persons, nil
}

func (r *PersonRepository) Delete(id string) error {
	return fmt.Errorf("err")
}

func (r PersonRepository) unmarshalPerson(
	id string,
	firstName string,
	lastName string,
	mail string,
	pwd string,
) (*person.Person, error) {
	var email person.Email
	if err := email.UnmarshalText(mail); err != nil {
		return nil, err
	}
	var name person.Name
	if err := name.UnmarshalText(firstName, lastName); err != nil {
		return nil, err
	}
	var password person.Password
	if err := password.UnmarshalText(pwd); err != nil {
		return nil, err
	}
	return person.UnmarshalPersonFromDatabase(id, email, name, password), nil
}
