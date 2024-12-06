package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"todo/configs"
	"todo/internal/database"
	"todo/internal/handlers"
	"todo/internal/service/task"
	"todo/internal/web/tasks"
)

func main() {
	conf := configs.LoadConfig()

	db := database.NewDB(conf)

	repo := task.NewTaskRepository(db)
	serv := task.NewTaskService(repo)

	handler := handlers.NewTaskHandler(serv)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	log.Fatal(e.Start(conf.Path.Port))

}
