package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	// writing tests in go programming language
	// Go's built-in support for unit testing makes
	// it easier to test as you go

	erfan, err := NewPerson("erfan", 17)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(erfan)
}

func NewPerson(name string, age int) (Person, error) {
	return Person{
		name, age,
	}, nil
}
