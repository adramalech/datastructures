package models

type Person struct {
	firstName string
	lastName string
	middleName string
}

func NewPerson(firstName string, middleName string, lastName string) Person {
	return Person{
		firstName:  firstName,
		lastName:   lastName,
		middleName: middleName,
	}
}


func (p *Person) FullName() string {
	return p.firstName + " " + p.middleName + " " + p.lastName
}