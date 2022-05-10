package main

func main() {
	// A panic typically means something went unexpectedly
	// wrong. Mostly we use it to fail fast on errors that
	// shouldn’t occur during normal operation, or that we
	// aren’t prepared to handle gracefully.

	// if the error is unrecoverable sometimes the
	// program simply can't continue
	// for this purpose there is the built-in function
	// Panic() that is effect creates a run time error
	// that will stop the program

	myPanic()


	// this is only an example but real program or library
	// should avoid panic()
}

func myPanic() {
	panic("we have the error :)")
}
