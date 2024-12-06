package handlers

import (
	"context"
	"todo/internal/service/task"
	"todo/internal/web/api"
)

type TaskHandler struct {
	Service *task.TaskService
}

func NewTaskHandler(service *task.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

func (handler *TaskHandler) GetTasks(ctx context.Context, request api.GetTasksRequestObject) (api.GetTasksResponseObject, error) {
	allTasks, err := handler.Service.GetAllTask()
	if err != nil {
		return nil, err
	}

	response := api.GetTasks200JSONResponse{}

	for _, tas := range allTasks {
		newTask := api.Task{
			Id:     tas.Id,
			Title:  tas.Title,
			IsDone: tas.IsDone,
		}
		response = append(response, newTask)
	}

	return response, nil
}

func (handler *TaskHandler) PostTasks(ctx context.Context, request api.PostTasksRequestObject) (api.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := &api.Task{
		Title:  taskRequest.Title,
		IsDone: taskRequest.IsDone,
	}
	createdTask, err := handler.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := api.PostTasks201JSONResponse{
		Id:     createdTask.Id,
		Title:  createdTask.Title,
		IsDone: createdTask.IsDone,
	}

	return response, nil
}

func (handler *TaskHandler) PatchTasksId(ctx context.Context, request api.PatchTasksIdRequestObject) (api.PatchTasksIdResponseObject, error) {
	taskRequest := request.Body
	id := request.Id

	updatedTask, err := handler.Service.UpdateTaskById(id, &api.Task{
		Title:  taskRequest.Title,
		IsDone: taskRequest.IsDone,
	})
	if err != nil {
		return nil, err
	}

	response := api.PatchTasksId201JSONResponse{
		Id:     updatedTask.Id,
		Title:  updatedTask.Title,
		IsDone: updatedTask.IsDone,
	}
	return response, nil
}

func (handler *TaskHandler) DeleteTasksId(ctx context.Context, request api.DeleteTasksIdRequestObject) (api.DeleteTasksIdResponseObject, error) {
	id := request.Id

	err := handler.Service.DeleteTaskById(id)
	if err != nil {
		return nil, err
	}

	response := api.DeleteTasksId204Response{}

	return response, err
}
