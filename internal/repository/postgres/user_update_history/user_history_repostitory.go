package user_update_history

import (
	"context"
	"users/internal/DTO/history"

	"users/internal/repository"
	"users/internal/repository/postgres/db"

	sq "github.com/Masterminds/squirrel"
)

const (
	tableName = "users_update_history"

	idColumn       = "id"
	userIdColumn   = "user_id"
	emailOldColumn = "email_old"
	emailNewColumn = "email_new"
	nameOldColumn  = "name_old"
	nameNewColumn  = "name_new"
	roleOldColumn  = "role_old"
	roleNewColumn  = "role_new"
)

type userUpdateHistoryRepo struct {
	db db.Client
}

func NewUserUpdateHistoryRepo(db db.Client) repository.UserHistoryRepository {
	return &userUpdateHistoryRepo{db: db}
}

func (u userUpdateHistoryRepo) Create(ctx context.Context, userHistory *history.UserUpdateHistory) error {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(
			userIdColumn,
			emailOldColumn,
			emailNewColumn,
			nameOldColumn,
			nameNewColumn,
			roleOldColumn,
			roleNewColumn).
		Values(
			userHistory.UserId,
			userHistory.EmailOld,
			userHistory.EmailNew,
			userHistory.NameOld,
			userHistory.NameNew,
			userHistory.RoleOld,
			userHistory.RoleNew)

	q, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	query := db.Query{
		Name:     "user_history_repository.Create",
		QueryRaw: q,
	}

	_, err = u.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
