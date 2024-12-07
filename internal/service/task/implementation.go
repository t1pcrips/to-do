package task

import (
	"context"
	"errors"
	"todo/internal/model"
	"todo/internal/repository"
	"todo/internal/service"
)

type TaskServiceImpl struct {
	TaskRepository repository.TaskRepository
	UserRepository repository.UserRepository
}

func NewTaskServiceImpl(tRepo repository.TaskRepository, uRepo repository.UserRepository) service.TaskService {
	return &TaskServiceImpl{
		TaskRepository: tRepo,
		UserRepository: uRepo,
	}
}

func (s *TaskServiceImpl) CreateTaskForUser(ctx context.Context, info *model.CreateTaskRequest) (int64, error) {
	err := s.checkUser(ctx, info.UserId)
	if err != nil {
		return 0, err
	}

	taskId, err := s.TaskRepository.CreateTaskForUser(ctx, info)
	if err != nil {
		return 0, err
	}

	return taskId, nil
}

func (s *TaskServiceImpl) GetAllTasksUser(ctx context.Context, userId int64) ([]model.Task, error) {
	err := s.checkUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	tasks, err := s.TaskRepository.GetAllTasksUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *TaskServiceImpl) UpdateTaskForUser(ctx context.Context, info *model.UpdateTaskRequest) error {
	err := s.checkUser(ctx, info.Id)
	if err != nil {
		return err
	}

	err = s.checkTask(ctx, info.UserId)
	if err != nil {
		return err
	}

	err = s.TaskRepository.UpdateTaskForUser(ctx, info)
	if err != nil {
		return err
	}

	return nil
}

func (s *TaskServiceImpl) DeleteTaskForUser(ctx context.Context, userId int64, taskId int64) error {
	err := s.checkUser(ctx, userId)
	if err != nil {
		return err
	}

	err = s.checkTask(ctx, taskId)
	if err != nil {
		return err
	}

	err = s.TaskRepository.DeleteTaskForUser(ctx, taskId)
	if err != nil {
		return err
	}

	return nil
}

func (s *TaskServiceImpl) checkUser(ctx context.Context, userId int64) error {
	exists, err := s.UserRepository.CheckUserExists(ctx, userId)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("user do not exists")
	}

	return err
}

func (s *TaskServiceImpl) checkTask(ctx context.Context, taskId int64) error {
	exists, err := s.TaskRepository.CheckTaskExists(ctx, taskId)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("task do not exists")
	}

	return err
}
