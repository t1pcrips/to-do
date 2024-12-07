package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"todo/internal/configs"
	"todo/internal/database"
	"todo/internal/handlers/task"
	"todo/internal/handlers/user"
	taskRepo "todo/internal/repository/task"
	userRepo "todo/internal/repository/user"
	taskService "todo/internal/service/task"
	userService "todo/internal/service/user"
	apiTasks "todo/internal/web/tasks"
	apiUsers "todo/internal/web/users"
)

func main() {
	conf := configs.LoadConfig()
	ctx := context.Background()

	db := database.NewDB(ctx, conf)

	userRepository := userRepo.NewUserRepositoryImpl(db)
	taskRepository := taskRepo.NewTaskRepositoryImpl(db)

	userServ := userService.NewUserServiceImpl(taskRepository, userRepository)
	taskServ := taskService.NewTaskServiceImpl(taskRepository, userRepository)

	userHandler := user.NewUserHandler(userServ)
	taskHandler := task.NewTaskHandler(taskServ)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strickTask := apiTasks.NewStrictHandler(taskHandler, nil)
	strickUser := apiUsers.NewStrictHandler(userHandler, nil)

	apiTasks.RegisterHandlers(e, strickTask)
	apiUsers.RegisterHandlers(e, strickUser)

	log.Fatal(e.Start(conf.Path.Port))
}
