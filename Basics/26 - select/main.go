package main

import (
	"fmt"
	"time"
)

func main() {
	//  The select statement lets a goroutine wait on
	// multiple communication operations.
	// A select blocks until one of its cases can run,
	// then it executes that case. It chooses one at
	// random if multiple are ready.
	chan1, chan2 := make(chan int), make(chan string)

	go NewRoutine(chan1, chan2)

	for {
		select {
		case msg := <-chan1:
			fmt.Println(msg)

		case msgEnd := <-chan2:
			fmt.Println(msgEnd)
			return
		}
	}
}

func NewRoutine(c1 chan int, c2 chan string) {
	for i := 0; i < 5; i++ {
		c1 <- i
	}

	time.Sleep(1 * time.Second)
	c2 <- "end"
}
