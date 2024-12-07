package task

import (
	"context"
	"todo/internal/converter"
	"todo/internal/handlers"
	"todo/internal/service"
	apiTasks "todo/internal/web/tasks"
)

type TaskHanlderImpl struct {
	ServiceTask service.TaskService
}

func NewTaskHandler(tServ service.TaskService) handlers.TaskHandler {
	return &TaskHanlderImpl{
		ServiceTask: tServ,
	}
}

func (h *TaskHanlderImpl) PostTasks(ctx context.Context, req apiTasks.PostTasksRequestObject) (apiTasks.PostTasksResponseObject, error) {
	info := converter.ToModelFromApiCreateTask(&req)

	taskId, err := h.ServiceTask.CreateTaskForUser(ctx, info)
	if err != nil {
		return nil, err
	}

	return apiTasks.PostTasks201JSONResponse{
		Id: &taskId,
	}, nil
}

func (h *TaskHanlderImpl) GetTasks(ctx context.Context, req apiTasks.GetTasksRequestObject) (apiTasks.GetTasksResponseObject, error) {
	tasks, err := h.ServiceTask.GetAllTasksUser(ctx, req.Params.UserId)
	if err != nil {
		return nil, err
	}

	response := apiTasks.GetTasks200JSONResponse{}

	for _, tas := range tasks {
		newTask := apiTasks.Task{
			Id:     &tas.Id,
			UserId: &tas.UserId,
			Title:  &tas.Title,
			IsDone: &tas.IsDone,
		}
		response = append(response, newTask)
	}

	return response, nil
}

func (h *TaskHanlderImpl) PatchTasks(ctx context.Context, req apiTasks.PatchTasksRequestObject) (apiTasks.PatchTasksResponseObject, error) {
	info := converter.ToModelFromApiUpdateTask(&req)

	err := h.ServiceTask.UpdateTaskForUser(ctx, info)
	if err != nil {
		return nil, err
	}

	return apiTasks.PatchTasks200JSONResponse{}, nil
}

func (h *TaskHanlderImpl) DeleteTasks(ctx context.Context, req apiTasks.DeleteTasksRequestObject) (apiTasks.DeleteTasksResponseObject, error) {
	err := h.ServiceTask.DeleteTaskForUser(ctx, req.Params.UserId, req.Params.Id)
	if err != nil {
		return nil, err
	}

	return apiTasks.DeleteTasks204Response{}, nil
}
