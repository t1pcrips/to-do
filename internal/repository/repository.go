package repository

import (
	"context"
	"todo/internal/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, info *model.CreateUserRequest) (int64, error)
	UpdateUser(ctx context.Context, info *model.UpdateUserRequest) error
	DeleteUser(ctx context.Context, userId int64) error
	CheckUserExists(ctx context.Context, userId int64) (bool, error)
}

type TaskRepository interface {
	CreateTaskForUser(ctx context.Context, info *model.CreateTaskRequest) (int64, error)
	GetAllTasksUser(ctx context.Context, userId int64) ([]model.Task, error)
	UpdateTaskForUser(ctx context.Context, info *model.UpdateTaskRequest) error
	DeleteTaskForUser(ctx context.Context, taskId int64) error
	CheckTaskExists(ctx context.Context, taskId int64) (bool, error)
}
