package main

import "fmt"

func main() {
	// fallthrough keyword is used in switch
	// statement in golang. This keyword is used
	// in switch case block. If the fallthrough
	// keyword is present in the case block, then
	// it will transfer control to the next case
	// even though the current case might have
	// matched.

	val := 8

	switch {
	case val < 10:
		fmt.Println("is less than 10")
		fallthrough

	case val < 20:
		fmt.Println("is less than 20")

	case val < 30:
		fmt.Println("is less than 30")

	default:
		fmt.Println("is greater than 30")
	}
}
