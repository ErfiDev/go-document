package main

import (
	"fmt"
	"sync"
)

func main() {
	// Sync package Mutex feature

	/*
		Mutex feature fixes the big issue called
		Race conditions(data racing).

		When two or more goroutines access the same
		memory location at least one of them is a write.
	*/

	// if we run the code below with -race
	// we can see the data race

	// var x int
	// go func() {
	// 	x++
	// }()
	// fmt.Println(x)

	// Fixing the problem with `Mutex`

	m := new(sync.Mutex)

	var x int

	go func() {
		m.Lock()
		x++
		m.Unlock()
	}()

	m.Lock()
	fmt.Println(x)
	m.Unlock()

	// if we run the code with -race
	// we see everything is ok.
}
