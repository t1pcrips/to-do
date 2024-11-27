package service

import (
	"gorm.io/gorm/clause"
	"todo/internal/database"
	"todo/internal/web/tasks"
)

type TaskRepository struct {
	Database *database.Db
}

func NewTaskRepository(db *database.Db) *TaskRepository {
	return &TaskRepository{
		Database: db,
	}
}

func (repo *TaskRepository) CreateTask(task *tasks.Task) (*tasks.Task, error) {
	result := repo.Database.DB.Create(task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

func (repo *TaskRepository) GetAllTasks() ([]tasks.Task, error) {
	var tasks []tasks.Task
	result := repo.Database.DB.Table("tasks").Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (repo *TaskRepository) UpdateTask(task *tasks.Task) (*tasks.Task, error) {
	result := repo.Database.DB.Table("tasks").Clauses(clause.Returning{}).Updates(task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

func (repo *TaskRepository) DeleteTask(id uint) error {
	result := repo.Database.DB.Table("tasks").Delete(&tasks.Task{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *TaskRepository) FindByIdTask(id uint) (*tasks.Task, error) {
	var task tasks.Task
	result := repo.Database.DB.Table("tasks").First(&task, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}
