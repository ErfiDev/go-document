package main

import "fmt"

func main() {
	// golang can make choices with data.
	// This data (variables) are used with a condition:
	// if statements start with a condition
	// A condition may be (x > 3), (y < 4), (weather = rain).
	// Only if a condition is true, code is executed.

	x := 3

	// Condition
	if x > 10 {
		fmt.Println("x is higher than 10")
	}
	// this condition is not working becuase
	// condition is false

	// but this condition is working becuase
	// consition is true
	if x == 3 {
		fmt.Println("condition working")
	}

	// You can execute a codeblock if a condition
	// is not true
	a := 10

	if a > 20 {
		fmt.Println("a variable higher than 20")
	} else {
		fmt.Println("a variable is lower than 20")
	}
}
