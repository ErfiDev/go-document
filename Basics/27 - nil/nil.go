package main

import "fmt"

type person struct {
	name string
}

func main() {
	// Are nil and null the same thing?
	// Coming from Java or other programming language
	// background you might think they are but they
	// partly are the same yet somewhat different.
	// nil is a predeclared identifier in Go that 
	// represents zero values for pointers, interfaces, 
	// channels, maps, slices and function types. nil being 
	// the zero value of pointers and interfaces, uninitialized 
	// pointers and interfaces will be nil.So does that mean a 
	//  string can’t be nil, yes you are right, a string doesn’t 
	// qualify to be either of the above-mentioned types 
	// and hence can’t be assigned nil.

	// a := nil
	// This will be a compile-time error use of untyped nilsince
	// nil is untyped and the only untyped value that doesn’t 
	// have default type and the compiler has no idea which 
	// type it has to assign to a.

	erfan := getPerson()
	fmt.Println(erfan)

	var nilSlice []string = nil
	fmt.Println(nilSlice)

	var nilMap map[string]string = nil
	fmt.Println(nilMap)
}

func getPerson() *person {
	return nil
}
