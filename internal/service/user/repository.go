package user

import (
	"todo/internal/database"
)

const (
	tableUsers = "users"
)

type UserRopsitory struct {
	Database *database.Db
}

func NewUserRopsitory(database *database.Db) *UserRopsitory {
	return &UserRopsitory{
		Database: database,
	}
}

func (repo *UserRopsitory) CreateUser(info *User) error {
	result := repo.Database.DB.Table(tableUsers).Create(info)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *UserRopsitory) GetAllUsers(userId int64) ([]User, error) {
	var users []User

	result := repo.Database.DB.Table(tableUsers).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (repo *UserRopsitory) UpdateUser(userId int64, info *User) error {
	result := repo.Database.DB.Table(tableUsers).Updates(info).Where("id = ?", userId)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *UserRopsitory) DeleteUser(userId int64) error {
	result := repo.Database.DB.Table(tableUsers).Delete(&User{}, userId)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
