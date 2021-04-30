package person

type Person struct {
	Id       string   `json:"id"`
	Name     Name     `json:"name"`
	Email    Email    `json:"email"`
	Password Password `json:"-"`
}

func NewPerson(id string, name Name, email Email, password Password) *Person {
	return &Person{
		Id:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}
}

func UnmarshalPersonFromDatabase(
	id string,
	email Email,
	name Name,
	password Password,
) *Person {
	return NewPerson(id, name, email, password)
}
