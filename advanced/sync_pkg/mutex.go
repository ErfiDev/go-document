package main

import (
	"fmt"
	"sync"
)

var i int

func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	i = i + 1
	m.Unlock()

	wg.Done()
}

func main() {
	// Sync package Mutex feature

	/*
		Mutex feature fixes the big issue called
		Race conditions(data racing).

		When two or more goroutines access the same
		memory location at least one of them is a write.
	*/

	wg := new(sync.WaitGroup)
	m := new(sync.Mutex)

	for j := 0; j < 100; j++ {
		wg.Add(1)
		go worker(wg, m)
	}

	wg.Wait()
	fmt.Println(i)
}
