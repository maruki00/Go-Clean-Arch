package entities

/*

type Person struct{
	Id int
	Name string
	Email string
	Age int
}


*/

type PersonEntity interface {
	SetId(idi int)
	GetId() int
	GetName() string
	GetEmail() string
	GetAge() int
}
