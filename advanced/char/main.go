package main

import "fmt"

func main() {
	// golang does not have any datatype of `char`

	// but we can use `byte or rune` to represent a single character

	// byte is used for repersent of ASCII characters
	// and rune is used for ALL UNICODE characters

	// to create either a byte or rune we use single quotes

	var E = 'E'

	fmt.Printf("the type of E is %T\n", E)
	fmt.Println(string(E))

	// while declaring byte we have to specify the type,
	// if we don't specify the type, then the default
	// type is meant as a rune

	var A byte = 'A'
	fmt.Printf("type of A = %T\n", A)
	fmt.Println(string(A))

	// rune = int32
	// byte = uint8
}
