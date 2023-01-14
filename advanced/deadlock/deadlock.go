package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Deadlock:
		occur when two or more threads are blocked,
		waiting for each other to release a response.
*/

type Value struct {
	mu sync.Mutex
	v  int
}

func sum(v1, v2 *Value, wg *sync.WaitGroup) {
	defer wg.Done()

	v1.mu.Lock()
	defer v1.mu.Unlock()

	time.Sleep(time.Second * 2)

	v2.mu.Lock()
	defer v2.mu.Unlock()

	fmt.Println("sum: ", v1.v+v2.v)
}

func main() {
	wg := new(sync.WaitGroup)

	v1, v2 := new(Value), new(Value)

	wg.Add(2)

	go sum(v1, v2, wg)
	go sum(v2, v1, wg)

	wg.Wait()
}

/*
	What happend?

	our first call to `sum` lock `a` and then attempts
	to lock `b` but in the meantime our second call to
	`sum` has locked `b` and has attempts to lock `a`.
	both goroutines wait infinitely on each other.
*/
