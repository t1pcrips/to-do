package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	port = "8080"
	task string
	rb   requestBody
)

type requestBody struct {
	Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello task")

	w.WriteHeader(200)
	fmt.Fprintf(w, "Hello, %s", task)
}

func CreateNewTask(w http.ResponseWriter, r *http.Request) {
	log.Println("New task create")

	json.NewDecoder(r.Body).Decode(&rb)
	task = rb.Message

	w.WriteHeader(200)
	fmt.Fprint(w, "new task created")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/task", CreateNewTask).Methods("POST")

	log.Println("Server starts")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
