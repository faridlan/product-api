package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/faridlan/product-api/helper"
	"github.com/faridlan/product-api/model/domain"
	"github.com/faridlan/product-api/model/web"
	"github.com/faridlan/product-api/repository"
)

type ProductServiceImpl struct {
	ProductRepo repository.ProductRepository
	DB          *sql.DB
}

func NewProductService(productRepo repository.ProductRepository, DB *sql.DB) ProductService {
	return &ProductServiceImpl{
		ProductRepo: productRepo,
		DB:          DB,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreate) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	epochTimeNow := time.Now().Unix()

	product := domain.Product{
		Name:      request.Name,
		Price:     request.Price,
		Quantity:  request.Quantity,
		CreatedAt: epochTimeNow,
	}
	product = service.ProductRepo.Save(ctx, tx, product)
	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdate) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepo.FindById(ctx, tx, request.Id)
	helper.PanicIfErr(err)

	epochTimeNow := time.Now().Unix()

	product.Id = request.Id
	product.Name = request.Name
	product.Price = request.Price
	product.Quantity = request.Quantity
	product.UpdatedAt = epochTimeNow

	product = service.ProductRepo.Update(ctx, tx, product)
	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepo.FindById(ctx, tx, productId)
	helper.PanicIfErr(err)

	service.ProductRepo.Update(ctx, tx, product)
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId int) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepo.FindById(ctx, tx, productId)
	helper.PanicIfErr(err)

	return helper.ToProductResponse(product)

}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepo.FindAll(ctx, tx)

	return helper.ToProductResponses(products)
}
