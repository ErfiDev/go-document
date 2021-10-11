package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {
	// Go makes it possible to recover from a panic, by
	// using the recover built-in function. A recover
	// can stop a panic from aborting the program and let
	//  it continue with execution instead.

	// An example of where this can be useful: a server
	// wouldn’t want to crash if one of the client
	// connections exhibits a critical error. Instead,
	// the server would want to close that connection and
	// continue serving other clients. In fact, this is
	// what Go’s net/http does by default for HTTP servers.

	defer func() {
		if r := recover(); r != nil {

			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	fmt.Println("After mayPanic()")
}
