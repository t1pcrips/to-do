package repository_converter

import (
	"todo/internal/model"
	"todo/internal/repository/repository_model"
)

func ModelToRepoCreateUser(info *model.CreateUserRequest) *repository_model.CreateUserRequest {
	return &repository_model.CreateUserRequest{
		Email:    info.Email,
		Password: info.Password,
	}
}

func ModelToRepoUpdateUser(info *model.UpdateUserRequest) *repository_model.UpdateUserRequest {
	return &repository_model.UpdateUserRequest{
		Id:    info.Id,
		Email: info.Email,
	}
}

func ModelToRepoCreateTask(info *model.CreateTaskRequest) *repository_model.CreateTaskRequest {
	return &repository_model.CreateTaskRequest{
		UserId: info.UserId,
		Title:  info.Title,
		IsDone: info.IsDone,
	}
}

func ModelToRepoUpdateTask(info *model.UpdateTaskRequest) *repository_model.UpdateTaskRequest {
	return &repository_model.UpdateTaskRequest{
		Id:     info.Id,
		Title:  info.Title,
		IsDone: info.IsDone,
	}
}
