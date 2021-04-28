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

//func ToPerson(id string, fN string, lN string, e string, p string, c time.Time, u time.Time) *Person {
//	var p Person
//	var pJson person
//	pJson.Created = c
//	pJson.Id = id
//	pJson.Updated = u
//	pJson.Email.value = e
//	pJson.Password.value = p
//	pJson.Name.firstName = fN
//	pJson.Name.lastName = lN
//	err = json.Unmarshal(byte[pJson], &p)

//}

//func (p Person) ToJSON() personJson {
//	email := p.Email.Value()
//	name := p.Name.FullName()
//	createdAt := p.Created.Format(utils.DateLayout)
//	updatedAt := p.Updated.Format(utils.DateLayout)
//	person := personJson{Id: p.Id, Email: email, Name: name, CreatedAt: createdAt, UpdatedAt: updatedAt}
//	return person
//}
