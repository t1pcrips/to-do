package main

import "gorm.io/gorm"

type Message struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
	gorm.Model
}
