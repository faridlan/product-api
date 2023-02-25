package service

import (
	"context"
	"database/sql"

	"github.com/faridlan/product-api/helper"
	"github.com/faridlan/product-api/model/domain"
	"github.com/faridlan/product-api/model/web"
	"github.com/faridlan/product-api/repository"
	"github.com/golang-jwt/jwt/v4"
)

type AuthServiceImpl struct {
	UserRepo repository.UserRepository
	DB       *sql.DB
}

func NewAuthService(userRepo repository.UserRepository, DB *sql.DB) AuthService {
	return &AuthServiceImpl{
		UserRepo: userRepo,
		DB:       DB,
	}
}

func (service *AuthServiceImpl) Login(ctx context.Context, request web.UserCreate) (web.LoginResponse, web.Claims) {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	strRandom := helper.RandStringRunes(20)

	user := domain.User{
		Username: request.Username,
		Password: request.Password,
	}

	user, err = service.UserRepo.Login(ctx, tx, user)
	if err != nil {
		panic(err)
	}

	claim := web.Claims{
		Id:       user.Id,
		Username: user.Username,
		Token:    strRandom,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(web.ExpiredTime),
		},
	}

	claimResult := helper.ToLoginResponse(user)

	return claimResult, claim
}
