package main

import "net/http"

// Basic Example of web servers
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	http.ListenAndServe(":3000", nil)
}
