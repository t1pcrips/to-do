package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"todo/internal/service"
	"todo/pkg/req"
	"todo/pkg/resp"
)

type TaskHandler struct {
	Service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

func (handler *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	task, err := req.HandleBody[service.Task](w, r)
	if err != nil {
		resp.Json(w, err.Error(), http.StatusBadRequest)
		return
	}
	task, err = handler.Service.CreateTask(task)
	if err != nil {
		resp.Json(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp.Json(w, task, http.StatusCreated)
}

func (handler *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := handler.Service.GetAllTask()
	if err != nil {
		resp.Json(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp.Json(w, tasks, http.StatusCreated)
}

func (handler *TaskHandler) UpdateTaskById(w http.ResponseWriter, r *http.Request) {

	task, err := req.HandleBody[service.Task](w, r)
	if err != nil {
		resp.Json(w, err.Error(), http.StatusBadRequest)
		return
	}
	idString := mux.Vars(r)["id"]
	task, err = handler.Service.UpdateTaskById(idString, task)
	if err != nil {
		resp.Json(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp.Json(w, task, http.StatusCreated)

}

func (handler *TaskHandler) DeleteTaskById(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	err := handler.Service.DeleteTaskById(idString)
	if err != nil {
		resp.Json(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp.Json(w, nil, http.StatusCreated)
}
