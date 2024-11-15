package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"net/http"
	"strconv"
)

var (
	port = ":8080"
)

func GetMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var messages []Message
	result := db.Table("messages").Find(&messages)

	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result.Error.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(messages)
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	newEntry := &Message{}
	json.NewDecoder(r.Body).Decode(newEntry)
	result := db.Table("messages").Create(newEntry)
	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result.Error.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idString := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	result := db.Table("messages").Delete(&Message{}, uint(id))
	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result.Error.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	newEntry := &Message{}
	json.NewDecoder(r.Body).Decode(newEntry)

	idString := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	message := &Message{
		Task:   newEntry.Task,
		IsDone: newEntry.IsDone,
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	result := db.Table("messages").Clauses(clause.Returning{}).Updates(message)
	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result.Error.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)
}

func main() {
	InitDB()
	err := db.AutoMigrate(Message{})
	if err != nil {
		log.Fatal("failed to migrate database")
	}

	router := mux.NewRouter()
	server := http.Server{
		Addr:    port,
		Handler: router,
	}

	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	router.HandleFunc("/api/message", CreateMessage).Methods("POST")
	router.HandleFunc("/api/message/{id}", UpdateMessage).Methods("PATCH")
	router.HandleFunc("/api/message/{id}", DeleteMessage).Methods("DELETE")

	log.Println("Server starts...")
	log.Fatal(server.ListenAndServe())
}
