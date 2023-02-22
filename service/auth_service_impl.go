package service

import (
	"context"
	"database/sql"

	"github.com/faridlan/product-api/helper"
	"github.com/faridlan/product-api/model/web"
	"github.com/faridlan/product-api/repository"
)

type AuthServiceImpl struct {
	UserRepo repository.UserRepository
	DB       *sql.DB
}

func (service *AuthServiceImpl) Login(ctx context.Context, request web.UserCreate) web.LoginResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	// user := domain.User{
	// 	Username: request.Username,
	// 	Password: request.Password,
	// }

	// u, err := service.UserRepo.Login(ctx, tx, user)

	return web.LoginResponse{}

}
