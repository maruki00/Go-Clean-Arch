package main

import (
	"fmt"

	"golang.org/x/tools/go/analysis/passes/stringintconv"
)

type Person struct {
	Id       int
	Name     string
	LastName string
}

func main() {

	fmt.Print("hello world")
}
