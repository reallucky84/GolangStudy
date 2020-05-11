package main

import (
	"fmt"
	"net/http"
)

type barHandler struct{}

func (b barHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello bar1.")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world!")
	})

	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is foo.")
	})

	http.Handle("/bar", &barHandler{})

	http.ListenAndServe(":3000", nil)
}
