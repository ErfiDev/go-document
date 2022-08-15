package main

import (
	"fmt"
	"sync"
)

func main() {
	// Sync package WaitGroup feature

	/*
		Golang Waitgroup allows you to block a
		specific code block to allow a set of
		goroutines to complete execution.
		An example would be to block the main
		function until the goroutines are completed
		and then unblocks the group.
	*/

	wg := new(sync.WaitGroup)

	// Adds delta, which mey be negative, if the counter
	// become zero all go routines blocked on Wait are
	// released, if counter goes negative Add panics
	wg.Add(1)

	go printSomthing("from another go routine", wg)

	fmt.Println("from main go routine")

	// Wait blocks until the WaitGroup counter is zero
	wg.Wait()
}

func printSomthing(s string, wg *sync.WaitGroup) {
	// Done decrements the WaitGroup counter by one
	defer wg.Done()

	fmt.Println(s)
}
