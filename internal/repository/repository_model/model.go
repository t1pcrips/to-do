package repository_model

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
	Title  string
	IsDone bool
}
