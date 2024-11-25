package models

import valueobject "go-clean-arch/internal/Person/Domain/ValueObjects"

type Person struct {
	Id    int
	Name  string
	Email *valueobject.Email
	Age   *valueobject.Age
}

func (obj *Person) GetId() int {
	return obj.Id
}
func (obj *Person) GetName() string {
	return obj.Name
}
func (obj *Person) GetEmail() string {
	return obj.Email.String()
}
func (obj *Person) GetAge() int {
	return obj.Age.GetValue()
}

func (obj *Person) SetId(id int) {
	obj.Id = id
}
