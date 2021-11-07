package main

import "net/http"

func main() {
	Router := http.NewServeMux()

	Router.HandleFunc("/", Home)
	Router.HandleFunc("/about", About)

	server := http.Server{
		Addr:    ":3000",
		Handler: Router,
	}

	server.ListenAndServe()
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("`/` Route"))
}

func About(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("`/about` Route"))
}
