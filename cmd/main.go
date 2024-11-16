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

	router.HandleFunc("/task", handler.CreateTask()).Methods("POST")
	router.HandleFunc("/task/{id}", handler.UpdateTask()).Methods("PATCH")
	router.HandleFunc("/task/{id}", handler.DeleteTask()).Methods("DELETE")
	router.HandleFunc("/tasks", handler.GetAllTasks()).Methods("GET")

	log.Println("Server starts...")
	log.Fatal(server.ListenAndServe())
}
