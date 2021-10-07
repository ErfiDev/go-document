package main

import (
	"log"
	"os"
)

func main() {
	// Defer is used to ensure that a function call is
	// performed later in a program’s execution, usually
	// for purposes of cleanup. defer is often used where
	// e.g. ensure and finally would be used in other
	// languages.

	// Suppose we wanted to create a file, write to it,
	// and then close when we’re done. Here’s how we could
	// do that with defer.

	// Immediately after getting a file object with
	// createFile, we defer the closing of that file
	// with closeFile. This will be executed at the end
	// of the enclosing function (main), after writeFile
	// has finished.

	file , err := os.Create("test.txt")
	defer file.Close()
	// if we remove the defer keyword we receive the big error
	// called : file already closed
	if err != nil {
		log.Fatal(err)
	}

	writeFile(file , "hi")
}

func writeFile(file *os.File , data string) {
	toByteSlice := []byte(data)
	_ , err := file.Write(toByteSlice)
	if err != nil {
		panic(err)
	}
}
