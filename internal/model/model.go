package model

import "time"

type User struct {
	Id        int64
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Task struct {
	Id        int64     `db:"id"`
	UserId    int64     `db:"user_id"`
	Title     string    `db:"title"`
	IsDone    bool      `db:"is_done"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type CreateUserRequest struct {
	Email    string
	Password string
}

type UpdateUserRequest struct {
	Id    int64
	Email string
}

type CreateTaskRequest struct {
	UserId int64
	Title  string
	IsDone bool
}

type UpdateTaskRequest struct {
	Id     int64
	UserId int64
	Title  string
	IsDone bool
}
