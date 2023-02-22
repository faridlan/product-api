package repository

import (
	"context"
	"database/sql"

	"github.com/faridlan/product-api/model/domain"
)

type UserRepository interface {
	Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
}
