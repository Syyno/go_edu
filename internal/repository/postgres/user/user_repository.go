package repository

import (
	"context"
	"time"
	models "users/internal/DTO/user"
	appErrors "users/internal/domain/errors"
	"users/internal/domain/user"
	"users/internal/repository"
	"users/internal/repository/postgres/converter"
	"users/internal/repository/postgres/db"
	"users/internal/repository/postgres/user/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

const (
	tableName = "users"

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	passwordColumn  = "password"
	roleColumn      = "role"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"

	entryNotFoundErrorMsg = "no rows in result set"
)

type userRepository struct {
	db db.Client
}

func NewUserRepository(db db.Client) repository.UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Get(ctx context.Context, id int64) (*user.UserDomain, error) {
	opertation := "user_repository.Get"
	builder := sq.Select(
		idColumn,
		nameColumn,
		emailColumn,
		passwordColumn,
		roleColumn,
		createdAtColumn,
		updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     opertation,
		QueryRaw: query,
	}

	var user model.UserRepo
	err = u.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		if err.Error() == entryNotFoundErrorMsg {
			return nil, &appErrors.UserNotFoundError{
				Id: id,
			}
		}
		return nil, err
	}

	return converter.ToDomain(&user), nil
}

func (u *userRepository) Create(ctx context.Context, userDomain user.UserDomain) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passwordColumn, roleColumn).
		Values(userDomain.Name, userDomain.Email, userDomain.Password, userDomain.Role).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = u.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *userRepository) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).
		Where(sq.Eq{idColumn: id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Delete",
		QueryRaw: query,
	}

	c, err := u.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	if c.RowsAffected() != 1 {
		return errors.New("Not found entry with provided id")
	}

	return nil
}

func (u *userRepository) Update(ctx context.Context, userUpdate models.UserUpdate) error {

	builder := sq.Update(tableName)

	if userUpdate.EmailProvided {
		builder = builder.Set(emailColumn, userUpdate.EmailValue)
	}

	if userUpdate.NameProvided {
		builder = builder.Set(nameColumn, userUpdate.NameValue)
	}

	if userUpdate.RoleProvided {
		builder = builder.Set(roleColumn, userUpdate.RoleValue)
	}

	builder = builder.
		Set(updatedAtColumn, time.Now().UTC()).
		Where(sq.Eq{idColumn: userUpdate.Id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	r, err := u.db.DB().ExecContext(ctx, q, args...)

	if err != nil {
		return err
	}

	if r.RowsAffected() != 1 {
		return errors.New("Not found entry with provided id")
	}

	return nil
}
