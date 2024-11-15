package service

import (
	"gorm.io/gorm/clause"
	"todo/internal/database"
)

type TaskRepository struct {
	Database *database.Db
}

func NewTaskRepository(db *database.Db) *TaskRepository {
	return &TaskRepository{
		Database: db,
	}
}

func (repo *TaskRepository) Create(task *Task) (*Task, error) {
	result := repo.Database.DB.Create(task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

func (repo *TaskRepository) GetAll() ([]Task, error) {
	var tasks []Task
	result := repo.Database.DB.Table("tasks").Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (repo *TaskRepository) Update(task *Task) (*Task, error) {
	result := repo.Database.DB.Table("tasks").Clauses(clause.Returning{}).Updates(task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

func (repo *TaskRepository) Delete(id uint) error {
	result := repo.Database.DB.Table("tasks").Delete(&Task{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *TaskRepository) FindById(id uint) (*Task, error) {
	var task Task
	result := repo.Database.DB.Table("tasks").First(&task, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}
