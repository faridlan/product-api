package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/product-api/helper"
	"github.com/faridlan/product-api/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {

	SQL := "SELECT id, username FROM user WHERE username = ? and password = ?"

	rows, err := tx.QueryContext(ctx, SQL, user.Username, user.Password)
	helper.PanicIfErr(err)

	defer rows.Close()

	if rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Username)
		helper.PanicIfErr(err)

		return user, nil
	} else {
		return user, errors.New("unauthorized")
	}
}
