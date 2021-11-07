package main

import (
	"html/template"
	"net/http"
)

func main() {
	// Initializing template
	tmp, _ := template.ParseFiles("./templates/home.gohtml")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Executing template
		tmp.Execute(w, map[string]string{
			"title": "home page title",
		})
	})

	http.ListenAndServe(":3000", nil)
}
