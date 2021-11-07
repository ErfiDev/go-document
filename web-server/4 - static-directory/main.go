package main

import "net/http"

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	// If we start the program and going on "/index.html" path
	// we see the index.html page on the static folder
	// that is easy!

	http.ListenAndServe(":3000", nil)
}
