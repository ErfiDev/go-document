package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Livelock:

		A livelock is like two friends trying to give
		each other a pen. Each friend has a pen and
		wants to give it to the other friend, but they
		both keep holding on to their own pen because
		they want the other person to give them their
		pen first. Neither of them can give their pen
		because they are both waiting for the other
		person to give their pen first. This is a l
		ivelock. In concurrent programming, livelocks
		happen when two or more tasks are unable to
		proceed because they are waiting for each other.
		Just like in the example, if the two friends had
		agreed to put their pens on a table and then pick
		up the pen they wanted, the livelock would be solved.
*/

func main() {
	var lock1, lock2 sync.Mutex

	go func() {
		lock1.Lock()
		fmt.Println("goroutine 1 lock 1")

		time.Sleep(time.Second)

		lock2.Lock()
		fmt.Println("goroutine 1 lock 2")

		lock2.Unlock()
		lock1.Unlock()
	}()

	go func() {
		lock2.Lock()
		fmt.Println("goroutine 2 lock 2")

		time.Sleep(time.Second)

		lock1.Lock()
		fmt.Println("goroutine 2 lock 1")

		lock1.Unlock()
		lock2.Unlock()
	}()

	time.Sleep(time.Second * 10)
}
