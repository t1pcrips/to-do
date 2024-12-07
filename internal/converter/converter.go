package converter

import (
	"todo/internal/model"
	apiTasks "todo/internal/web/tasks"
	apiUser "todo/internal/web/users"
)

func ToModelFromApiCreateUser(req *apiUser.PostUsersRequestObject) *model.CreateUserRequest {
	return &model.CreateUserRequest{
		Email:    *req.Body.Email,
		Password: *req.Body.Password,
	}
}

func ToModelFromApiUpdateUser(req *apiUser.PatchUsersRequestObject) *model.UpdateUserRequest {
	return &model.UpdateUserRequest{
		Id:    req.Params.Id,
		Email: *req.Body.Email,
	}
}

func ToModelFromApiCreateTask(req *apiTasks.PostTasksRequestObject) *model.CreateTaskRequest {
	return &model.CreateTaskRequest{
		UserId: *req.Body.UserId,
		Title:  *req.Body.Title,
		IsDone: *req.Body.IsDone,
	}
}

func ToModelFromApiUpdateTask(req *apiTasks.PatchTasksRequestObject) *model.UpdateTaskRequest {
	return &model.UpdateTaskRequest{
		Id:     req.Params.Id,
		UserId: *req.Body.UserId,
		Title:  *req.Body.Title,
		IsDone: *req.Body.IsDone,
	}
}
