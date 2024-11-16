package valueobject

import "fmt"

type Age struct {
	age int
}

func NewAge(age int) *Age {
	if age < 0 {
		panic("age should be more than 0")
	}
	return &Age{
		age: age,
	}
}

func (obj *Age) String() string {
	return fmt.Sprintf("%d", obj.age)
}

func (obj *Age) GetValue() int {
	return obj.age
}
