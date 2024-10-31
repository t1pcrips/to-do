package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	port = "8080"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
