package main

import "fmt"

type Msg struct {
	Desc string
	From string
}

func main() {
	// The new built-in function allocates memory.
	// The first argument is a type, not a value,
	// and the value returned is a pointer to a newly
	// allocated zero value of that type.

	// new(Type)
	newMsg := new(Msg)
	fmt.Printf("%#v\n", newMsg)
	Change(newMsg)
	fmt.Printf("%#v\n", newMsg)

	// The make built-in function allocates and
	// initializes an object of type slice, map,
	// or chan (only). Like new, the first argument
	// is a type, not a value. Unlike new, make’s
	// return type is the same as the type of its
	// argument, not a pointer to it.

	// make(Type, Size)
	var makeSomthing = make([]string, 2)
	fmt.Printf("%#v\n", makeSomthing)
}

func Change(msg *Msg) {
	msg.Desc = "Changed description"
}
