package main

import (
	"fmt"
	"errors"
)

// the blank identifier

func main() {
	// The use of a blank identifier in a for range
	// loop is a special case of a general situation:
	// multiple assignment.
	// If an assignment requires multiple values on the
	// left side, but one of the values will not be used
	// by the program, a blank identifier on the
	// left-hand-side of the assignment avoids the need
	// to create a dummy variable and makes it clear
	// that the value is to be discarded. For instance,
	// when calling a function that returns a value and
	// an error, but only the error is important, use
	// the blank identifier to discard the irrelevant value.

	// Example 1
	users := []string{"erfan", "john" , "tom" , "dominic"}

	for _ , value := range users {
		fmt.Println(value)
	}

	// Example 2
	_ , err := ExampleTwo(19)
	if err != nil {
		fmt.Println(err)
	}
}

func ExampleTwo(i int) (string , interface{}) {
	if i < 18 {
		return "" , errors.New("error")
	} else {
		return "ok" , ""
	}
}
