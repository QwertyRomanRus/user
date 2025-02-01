package user

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	model "user/internal/model"
	"user/internal/repository/user/convertor"
	modelRepo "user/internal/repository/user/model"
)

func (r *repo) Get(ctx context.Context, id int) (*model.User, error) {
	sql, args, err := sq.
		Select(idColumn, firstNameColumn, lastNameColumn, emailColumn, ageColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id}).
		Limit(1).
		ToSql()

	if err != nil {
		return nil, err
	}

	var userRepo modelRepo.User

	err = r.db.QueryRow(ctx, sql, args...).
		Scan(&userRepo.ID, &userRepo.Info.FirstName, &userRepo.Info.LastName, &userRepo.Info.Email, &userRepo.Info.Age)
	if err != nil {
		return nil, err
	}

	return convertor.ToUserFromRepo(&userRepo), nil
}
