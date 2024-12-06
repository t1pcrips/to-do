package user

import (
	"gorm.io/gorm"
	"todo/internal/web/api"
)

func FromApiToModel(info *api.User) *User {
	return &User{
		Email:    *info.Email,
		Password: *info.Password,
		Model:    gorm.Model{},
	}
}

func FromModelToApi(info *User) *api.User {
	return &api.User{
		Email:    &info.Email,
		Password: &info.Password,
	}
}
