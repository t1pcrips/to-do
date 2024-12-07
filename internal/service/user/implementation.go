package user

import (
	"context"
	"errors"
	"todo/internal/model"
	"todo/internal/repository"
	"todo/internal/service"
)

type UserServiceImpl struct {
	TaskRepository repository.TaskRepository
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(tRepo repository.TaskRepository, uRepo repository.UserRepository) service.UserService {
	return &UserServiceImpl{
		TaskRepository: tRepo,
		UserRepository: uRepo,
	}
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, info *model.CreateUserRequest) (int64, error) {
	userId, err := s.UserRepository.CreateUser(ctx, info)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (s *UserServiceImpl) UpdateUser(ctx context.Context, info *model.UpdateUserRequest) error {
	err := s.checkUser(ctx, info.Id)
	if err != nil {
		return err
	}

	err = s.UserRepository.UpdateUser(ctx, info)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserServiceImpl) DeleteUser(ctx context.Context, userId int64) error {
	err := s.checkUser(ctx, userId)
	if err != nil {
		return err
	}

	err = s.UserRepository.DeleteUser(ctx, userId)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserServiceImpl) checkUser(ctx context.Context, userId int64) error {
	exists, err := s.UserRepository.CheckUserExists(ctx, userId)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("user do not exists")
	}

	return err
}
