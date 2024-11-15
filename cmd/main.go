package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"todo/configs"
	"todo/internal/database"
	"todo/internal/handlers"
	"todo/internal/service"
)

func main() {
	conf := configs.LoadConfig()

	db := database.NewDB(conf)
	db.AutoMigrate(&service.Task{})

	repo := service.NewTaskRepository(db)
	serv := service.NewTaskService(repo)

	handler := handlers.NewTaskHandler(serv)

	router := mux.NewRouter()
	server := http.Server{
		Addr:    conf.Path.Port,
		Handler: router,
	}

	router.HandleFunc("/task", handler.Create()).Methods("POST")
	router.HandleFunc("/task/{id}", handler.Update()).Methods("PATCH")
	router.HandleFunc("/task/{id}", handler.Delete()).Methods("DELETE")
	router.HandleFunc("/tasks", handler.GetAll()).Methods("GET")

	log.Println("Server starts...")
	log.Fatal(server.ListenAndServe())
}
