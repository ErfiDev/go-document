package main

import (
	"fmt"
	"sync"
	"time"
)

var philosophers = []string{"a", "b", "c", "d", "e"}
var wg sync.WaitGroup

func main() {
	wg.Add(len(philosophers))

	leftHand := new(sync.Mutex)

	for _, ph := range philosophers {
		rightHand := new(sync.Mutex)

		go Eat(ph, rightHand, leftHand)

		leftHand = rightHand
	}

	wg.Wait()
}

func Eat(name string, rightHand, leftHand *sync.Mutex) {
	defer wg.Done()

	fmt.Printf("%s seate\n", name)

	rightHand.Lock()
	leftHand.Lock()

	fmt.Printf("%s is eating\n", name)
	time.Sleep(time.Duration(2) * time.Second)

	rightHand.Unlock()
	leftHand.Unlock()

	fmt.Printf("%s has left the table!\n", name)
}
