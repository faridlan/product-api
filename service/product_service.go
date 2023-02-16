package service

import (
	"context"

	"github.com/faridlan/product-api/model/web"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreate) web.ProductResponse
	Update(ctx context.Context, request web.ProductUpdate) web.ProductResponse
	Delete(ctx context.Context, productId int)
	FindById(ctx context.Context, productId int) web.ProductResponse
	FindAll(ctx context.Context) []web.ProductResponse
	CreateMany(ctx context.Context, request []web.ProductCreate) []web.ProductResponse
	DeleteAll(ctx context.Context)
}
