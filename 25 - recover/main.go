package main

import "fmt"

func main() {
	// Go makes it possible to recover from a panic,
	// by using the recover built-in function. A recover
	// can stop a panic from aborting the program and
	// let it continue with execution instead.


	defer func () {
		if r := recover() ; r != nil {
			fmt.Printf("this error: %s" , r)
		}
	}()

	my()
}

func my(){
	panic("test panic")
}
