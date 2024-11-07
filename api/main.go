package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	port = "8080"
)

func GetMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Getting messages handler...")

	var messages []Message
	result := db.Find(&messages)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("failed to find messages: ", result.Error)
		w.Write([]byte("failed to find messages"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(messages)

	log.Println("Messages received")
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating new message handler...")

	newEntry := &Message{}
	json.NewDecoder(r.Body).Decode(newEntry)

	result := db.Create(newEntry)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("failed to created message: ", result.Error)
		w.Write([]byte("failed to created message"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	log.Println("New message created")
}

func main() {

	InitDB()

	err := db.AutoMigrate(Message{})
	if err != nil {
		log.Fatal("failed to migrate database")
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")

	log.Println("Server starts...")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
