package task

import (
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"time"
	"todo/internal/database"
	"todo/internal/model"
	"todo/internal/repository"
	"todo/internal/repository/repository_converter"
)

const (
	tableTask       = "tasks"
	idColumn        = "id"
	userIdColumn    = "user_id"
	titleColumn     = "title"
	isDoneColumn    = "is_done"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
	returningId     = "RETURNING id"
)

type TaskRepositoryImpl struct {
	Db *database.Db
}

func NewTaskRepositoryImpl(db *database.Db) repository.TaskRepository {
	return &TaskRepositoryImpl{
		Db: db,
	}
}

func (repo *TaskRepositoryImpl) CreateTaskForUser(ctx context.Context, info *model.CreateTaskRequest) (int64, error) {
	repoInfo := repository_converter.ModelToRepoCreateTask(info)

	builderCreateUser := squirrel.Insert(tableTask).
		PlaceholderFormat(squirrel.Dollar).
		Columns(userIdColumn, titleColumn, createdAtColumn).
		Values(repoInfo.UserId, repoInfo.Title, time.Now()).
		Suffix(returningId)

	query, args, err := builderCreateUser.ToSql()
	if err != nil {
		return 0, err
	}

	var taskId int64

	err = repo.Db.Pool.QueryRow(ctx, query, args...).Scan(&taskId)
	if err != nil {
		return 0, err
	}

	return taskId, nil
}

func (repo *TaskRepositoryImpl) GetAllTasksUser(ctx context.Context, userId int64) ([]model.Task, error) {
	builderGetTasksUser := squirrel.Select(idColumn, userIdColumn, titleColumn, isDoneColumn, createdAtColumn, updatedAtColumn).
		From(tableTask).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{userIdColumn: userId})

	query, args, err := builderGetTasksUser.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := repo.Db.Pool.Query(ctx, query, args)

	var tasks []model.Task

	err = pgxscan.ScanAll(&tasks, rows)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (repo *TaskRepositoryImpl) UpdateTaskForUser(ctx context.Context, info *model.UpdateTaskRequest) error {
	repoInfo := repository_converter.ModelToRepoUpdateTask(info)

	builderUpdateTask := squirrel.Update(tableTask).
		PlaceholderFormat(squirrel.Dollar).
		Set(titleColumn, repoInfo.Title).
		Set(isDoneColumn, repoInfo.IsDone).
		Set(updatedAtColumn, time.Now()).
		Where(squirrel.Eq{idColumn: repoInfo.Id})

	query, args, err := builderUpdateTask.ToSql()
	if err != nil {
		return err
	}

	result, err := repo.Db.Pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("no rows to update")
	}

	return nil
}

func (repo *TaskRepositoryImpl) DeleteTaskForUser(ctx context.Context, taskId int64) error {
	builderDeleteTask := squirrel.Delete(tableTask).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{idColumn: taskId})

	query, args, err := builderDeleteTask.ToSql()
	if err != nil {
		return err
	}

	result, err := repo.Db.Pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("no rows to delete")
	}

	return nil
}

func (repo *TaskRepositoryImpl) CheckTaskExists(ctx context.Context, taskId int64) (bool, error) {
	builderCheckTask := squirrel.Select("1").
		From(tableTask).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{idColumn: taskId})

	query, args, err := builderCheckTask.ToSql()
	if err != nil {
		return false, err
	}

	var exists string

	err = repo.Db.Pool.QueryRow(ctx, query, args...).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists == "1", nil
}
