package task

import "gorm.io/gorm"

type Task struct {
	Text   string `json:"text"`
	IsDone bool   `json:"isDone"`
	gorm.Model
}
