package main

import "fmt"

type human struct {
	name string
}

type animal struct {
	name string
}

type car struct {
	name string
}

type testInterface interface {
	// types contain Speak() method
	Speak() string
}

func main() {
	// Interfaces are named collections of method signatures.
	//For our example we’ll implement this interface on human and animal types.
	erfan := human{"erfan"}
	dog := animal{"pitter"}
	dodge := car{"dodge raptor"}
	// human and animal type contain Speak()
	// method so can access to log() function

	log(erfan)
	log(dog)

	// log(dodge)
	// Error: cannot use dodge (variable of type car)
	// as testInterface value in argument to log:
	// missing method Speak()
	fmt.Println(dodge)
}

func log(in testInterface) {
	fmt.Println(in.Speak())
}

func (human) Speak() string {
	return "hellooooo"
}

func (animal) Speak() string {
	return "hop hop"
}
