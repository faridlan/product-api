package service

import (
	"context"

	"github.com/faridlan/product-api/model/web"
)

type AuthService interface {
	Login(ctx context.Context, request web.UserCreate) (web.LoginResponse, web.Claims)
}
