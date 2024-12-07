package handlers

import (
	"context"
	apiTasks "todo/internal/web/tasks"
	apiUsers "todo/internal/web/users"
)

// UserHandler описывает методы для обработки запросов, связанных с пользователями.
type UserHandler interface {
	PostUsers(ctx context.Context, req apiUsers.PostUsersRequestObject) (apiUsers.PostUsersResponseObject, error)
	PatchUsers(ctx context.Context, req apiUsers.PatchUsersRequestObject) (apiUsers.PatchUsersResponseObject, error)
	DeleteUsers(ctx context.Context, req apiUsers.DeleteUsersRequestObject) (apiUsers.DeleteUsersResponseObject, error)
}

// TaskHandler описывает методы для обработки запросов, связанных с задачами.
type TaskHandler interface {
	PostTasks(ctx context.Context, req apiTasks.PostTasksRequestObject) (apiTasks.PostTasksResponseObject, error)
	GetTasks(ctx context.Context, req apiTasks.GetTasksRequestObject) (apiTasks.GetTasksResponseObject, error)
	PatchTasks(ctx context.Context, req apiTasks.PatchTasksRequestObject) (apiTasks.PatchTasksResponseObject, error)
	DeleteTasks(ctx context.Context, req apiTasks.DeleteTasksRequestObject) (apiTasks.DeleteTasksResponseObject, error)
}
