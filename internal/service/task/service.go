package task

import (
	"todo/internal/web/api"
)

type TaskService struct {
	Repository *TaskRepository
}

func NewTaskService(repo *TaskRepository) *TaskService {
	return &TaskService{Repository: repo}
}

func (s *TaskService) GetAllTask() ([]api.Task, error) {
	return s.Repository.GetAllTasks()
}

func (s *TaskService) CreateTask(task *api.Task) (*api.Task, error) {
	return s.Repository.CreateTask(task)
}

func (s *TaskService) DeleteTaskById(id uint) error {
	_, err := s.Repository.FindByIdTask(id)
	if err != nil {
		return err
	}
	return s.Repository.DeleteTask(id)
}

func (s *TaskService) UpdateTaskById(id uint, task *api.Task) (*api.Task, error) {
	_, err := s.Repository.FindByIdTask(id)
	if err != nil {
		return nil, err
	}
	newTask := &api.Task{
		Id:     &id,
		Title:  task.Title,
		IsDone: task.IsDone,
	}
	return s.Repository.UpdateTask(newTask)
}
