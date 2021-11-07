package main

import (
	"fmt"
	"net/http"
)

func main() {
	Router := http.NewServeMux()

	Router.HandleFunc("/", Home)
	// Router → Middleware Handler → Application Handler
	Router.Handle("/about", Middleware(About))

	server := http.Server{
		Addr:    ":3000",
		Handler: Router,
	}

	server.ListenAndServe()
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("`/` Route"))
}

func Middleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware mission success")
		next.ServeHTTP(w, r)
	})
}

func About(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("`/about` Route with middleware"))
}
