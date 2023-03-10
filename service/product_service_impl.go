package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/faridlan/product-api/exception"
	"github.com/faridlan/product-api/helper"
	"github.com/faridlan/product-api/helper/mysql"
	"github.com/faridlan/product-api/model/domain"
	"github.com/faridlan/product-api/model/web"
	"github.com/faridlan/product-api/repository"
	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepo repository.ProductRepository
	DB          *sql.DB
	Validate    *validator.Validate
}

func NewProductService(productRepo repository.ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepo: productRepo,
		DB:          DB,
		Validate:    validate,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreate) web.ProductResponse {
	err := service.Validate.Struct(request)
	errs := helper.TranslateError(err, service.Validate)

	if err != nil {
		panic(exception.NewValidationError(errs))
	}

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	epochTimeNow := time.Now().UnixMilli()

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
	err := service.Validate.Struct(request)
	errs := helper.TranslateError(err, service.Validate)

	if err != nil {
		panic(exception.NewValidationError(errs))
	}

	for _, err := range errs {
		fmt.Println(err)
		return web.ProductResponse{}
	}

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepo.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewInterfaceError(err.Error()))
	}

	epochTimeNow := time.Now().UnixMilli()

	NewInt := mysql.NewNullInt64(epochTimeNow)

	product.Id = request.Id
	product.Name = request.Name
	product.Price = request.Price
	product.Quantity = request.Quantity
	product.UpdatedAt = NewInt

	product = service.ProductRepo.Update(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepo.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewInterfaceError(err.Error()))
	}

	service.ProductRepo.Delete(ctx, tx, product)
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId int) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepo.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewInterfaceError(err.Error()))
	}

	return helper.ToProductResponse(product)

}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepo.FindAll(ctx, tx)

	return helper.ToProductResponses(products)
}

func (service *ProductServiceImpl) CreateMany(ctx context.Context, request []web.ProductCreate) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	products := []domain.Product{}

	for _, req := range request {
		product := domain.Product{}

		product.Name = req.Name
		product.Price = req.Price
		product.Quantity = req.Quantity
		product.CreatedAt = time.Now().UnixMilli()

		products = append(products, product)
	}

	service.ProductRepo.SaveMany(ctx, tx, products)

	log.Println("Service", request)

	return helper.ToProductResponses(products)
}

func (service *ProductServiceImpl) DeleteAll(ctx context.Context) {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	service.ProductRepo.DeleteAll(ctx, tx)
}
