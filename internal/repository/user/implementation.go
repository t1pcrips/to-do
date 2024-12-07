package user

import (
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"time"
	"todo/internal/database"
	"todo/internal/model"
	"todo/internal/repository"
	"todo/internal/repository/repository_converter"
)

const (
	tableUser       = "users"
	idColumn        = "id"
	emailColumn     = "email"
	passwordColumn  = "password"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
	returningId     = "RETURNING id"
)

type UserRepositoryImpl struct {
	Db *database.Db
}

func NewUserRepositoryImpl(db *database.Db) repository.UserRepository {
	return &UserRepositoryImpl{
		Db: db,
	}
}

func (repo *UserRepositoryImpl) CreateUser(ctx context.Context, info *model.CreateUserRequest) (int64, error) {
	repoInfo := repository_converter.ModelToRepoCreateUser(info)

	builderCreateUser := squirrel.Insert(tableUser).
		PlaceholderFormat(squirrel.Dollar).
		Columns(emailColumn, passwordColumn, createdAtColumn).
		Values(repoInfo.Email, repoInfo.Password, time.Now()).
		Suffix(returningId)

	query, args, err := builderCreateUser.ToSql()
	if err != nil {
		return 0, err
	}

	var userId int64

	err = repo.Db.Pool.QueryRow(ctx, query, args...).Scan(&userId)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (repo *UserRepositoryImpl) UpdateUser(ctx context.Context, info *model.UpdateUserRequest) error {
	repoInfo := repository_converter.ModelToRepoUpdateUser(info)

	builderUpdateUser := squirrel.Update(tableUser).
		PlaceholderFormat(squirrel.Dollar).
		Set(emailColumn, repoInfo.Email).
		Set(updatedAtColumn, time.Now()).
		Where(squirrel.Eq{idColumn: repoInfo.Id})

	query, args, err := builderUpdateUser.ToSql()
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

func (repo *UserRepositoryImpl) DeleteUser(ctx context.Context, userId int64) error {
	builderDeleteUser := squirrel.Delete(tableUser).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{idColumn: userId})

	query, args, err := builderDeleteUser.ToSql()
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

func (repo *UserRepositoryImpl) CheckUserExists(ctx context.Context, userId int64) (bool, error) {
	builderCheckUser := squirrel.Select("1").
		From(tableUser).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{idColumn: userId})

	query, args, err := builderCheckUser.ToSql()
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
