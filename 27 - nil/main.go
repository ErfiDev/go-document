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
	// Coming from Java or other programming language
	// background you might think they are but they
	// partly are the same yet somewhat different.

	// a := nil
	// Coming from Java or other programming language
	// background you might think they are but they
	// partly are the same yet somewhat different.

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
