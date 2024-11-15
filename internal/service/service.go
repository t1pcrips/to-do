package service

import (
	"gorm.io/gorm"
	"strconv"
)

type TaskService struct {
	Repository *TaskRepository
}

func NewTaskService(repo *TaskRepository) *TaskService {
	return &TaskService{Repository: repo}
}

func (s *TaskService) GetAllTask() ([]Task, error) {
	return s.Repository.GetAll()
}

func (s *TaskService) CreateTask(task *Task) (*Task, error) {
	return s.Repository.Create(task)
}

func (s *TaskService) DeleteTask(idString string) error {
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		return err
	}
	uintID := uint(id)
	_, err = s.Repository.FindById(uintID)
	if err != nil {
		return err
	}
	return s.Repository.Delete(uintID)
}

func (s *TaskService) UpdateTask(idString string, task *Task) (*Task, error) {
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		return nil, err
	}
	uintID := uint(id)
	_, err = s.Repository.FindById(uintID)
	if err != nil {
		return nil, err
	}
	newTask := &Task{
		Text:   task.Text,
		IsDone: task.IsDone,
		Model: gorm.Model{
			ID: uintID,
		},
	}
	return s.Repository.Update(newTask)
}
