package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", hello).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>Gorilla Mux Router</h1>")
}
