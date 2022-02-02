package main

import "fmt"

func main() {
	// the goto keyword
	// for jumping to a labeled statement

	// how it work

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		if i == 5 {
			goto LOG
			// enter directly to the label
		}
		continue
	}

LOG:
	fmt.Println("LOG label print")
}
