package handlers

import (
	"context"
	"todo/internal/service"
	"todo/internal/web/tasks"
)

type TaskHandler struct {
	Service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

func (handler *TaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := handler.Service.GetAllTask()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tas := range allTasks {
		newTask := tasks.Task{
			Id:     tas.Id,
			Title:  tas.Title,
			IsDone: tas.IsDone,
		}
		response = append(response, newTask)
	}

	return response, nil
}

func (handler *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := &tasks.Task{
		Title:  taskRequest.Title,
		IsDone: taskRequest.IsDone,
	}
	createdTask, err := handler.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     createdTask.Id,
		Title:  createdTask.Title,
		IsDone: createdTask.IsDone,
	}

	return response, nil
}

func (handler *TaskHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskRequest := request.Body
	id := request.Id

	updatedTask, err := handler.Service.UpdateTaskById(id, &tasks.Task{
		Title:  taskRequest.Title,
		IsDone: taskRequest.IsDone,
	})
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId201JSONResponse{
		Id:     updatedTask.Id,
		Title:  updatedTask.Title,
		IsDone: updatedTask.IsDone,
	}
	return response, nil
}

func (handler *TaskHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := request.Id

	err := handler.Service.DeleteTaskById(id)
	if err != nil {
		return nil, err
	}

	response := tasks.DeleteTasksId204Response{}

	return response, err
}
