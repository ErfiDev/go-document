package main

import "testing"

// Now that you've gotten your code to a stable place
// (nicely done, by the way), add a test. Testing your
// code during development can expose bugs that find
// their way in as you make changes

func TestMain(t *testing.T) {
	_, err := NewPerson("test", 1)

	if err != nil {
		t.Error("test failed!")
	}
	// for run tests we should write
	// go test .  =>>  or  =>> go test -v  or  =>> go test

}
