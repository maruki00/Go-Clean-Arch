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
	GetId() int
	GetName() string
	GetEmail() string
	GetAge() int
}
