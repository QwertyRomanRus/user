package user

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	model "user/internal/model"
	"user/internal/repository/user/convertor"
	modelRepo "user/internal/repository/user/model"
)

// todo прокинуть limit из request
func (r *repo) List(ctx context.Context) ([]*model.User, error) {
	sql, args, err := sq.
		Select(idColumn, firstNameColumn, lastNameColumn, emailColumn, ageColumn).
		From(tableName).
		OrderBy("id ASC").
		Limit(limit).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(ctx, sql, args...)
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var userRepo modelRepo.User

		err = rows.Scan(&userRepo.ID, &userRepo.Info.FirstName, &userRepo.Info.LastName, &userRepo.Info.Email, &userRepo.Info.Age)
		if err != nil {
			return nil, err
		}

		users = append(users, convertor.ToUserFromRepo(&userRepo))
	}

	return users, nil
}
