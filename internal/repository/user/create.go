package user

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	model "user/internal/model"
	"user/internal/repository/user/convertor"
	modelRepo "user/internal/repository/user/model"
)

func (r *repo) Create(ctx context.Context, userInfo *model.UserInfo) (*model.User, error) {
	sql, args, err := sq.
		Insert(tableName).
		Columns(firstNameColumn, lastNameColumn, emailColumn, ageColumn).
		PlaceholderFormat(sq.Dollar).
		Values(userInfo.FirstName, userInfo.LastName, userInfo.Email, userInfo.Age).
		Suffix("RETURNING " + idColumn).
		ToSql()

	if err != nil {
		return nil, err
	}

	var user modelRepo.User
	user.Info = modelRepo.UserInfo(*userInfo)

	err = r.db.QueryRow(ctx, sql, args...).Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return convertor.ToUserFromRepo(&user), nil
}
