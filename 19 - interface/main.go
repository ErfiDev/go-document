package main

import "fmt"

type all interface {
	logger() string
}
type human struct{}
type animal struct{}
type user []string

func main() {
	h := human{}
	a := animal{}
	newUser := user{}

	printer(h)
	printer(a)
	printer(newUser)
}

func printer(a all) {
	fmt.Println(a.logger())
}

func (human) logger() string {
	return "hello human"
}

func (animal) logger() string {
	return "hello animal"
}

func (user) logger() string {
	return "hello user"
}
