package service_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/faridlan/product-api/helper"
	"github.com/faridlan/product-api/model/domain"
	"github.com/faridlan/product-api/model/web"
	"github.com/faridlan/product-api/repository"
	"github.com/faridlan/product-api/service"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func sqlMock() (*sql.DB, sqlmock.Sqlmock, error) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		return nil, nil, err
	} else {
		return db, mock, nil
	}

}

func NewService(db *sql.DB, validate *validator.Validate) service.ProductService {
	repository := repository.NewProductRepository()
	return service.NewProductService(
		repository,
		db,
		validate,
	)
}

func TestFindAll(t *testing.T) {

	db, mock, err := sqlMock()
	helper.PanicIfErr(err)
	validate := validator.New()

	ProductService := NewService(db, validate)

	product1 := domain.Product{
		Id:        1,
		Name:      "Laptop Asus",
		Price:     6000000,
		Quantity:  10,
		CreatedAt: time.Now().UnixMilli(),
		// UpdatedAt: &mysql.NullInt{},
	}

	product2 := domain.Product{
		Id:        2,
		Name:      "Laptop Lenovo",
		Price:     9000000,
		Quantity:  10,
		CreatedAt: time.Now().UnixMilli(),
		// UpdatedAt: &mysql.NullInt{},
	}

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT id, name, price, quantity, created_at, updated_at FROM product").WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"name",
		"price",
		"quantity",
		"created_at",
		"updated_at",
	}).AddRow(product1.Id, product1.Name, product1.Price, product1.Quantity, product1.CreatedAt, product1.UpdatedAt).AddRow(product2.Id, product2.Name, product2.Price, product2.Quantity, product2.CreatedAt, product2.UpdatedAt))
	mock.ExpectCommit()

	products := ProductService.FindAll(context.Background())

	productResponses := []web.ProductResponse{
		{
			Id:        product1.Id,
			Name:      product1.Name,
			Price:     product1.Price,
			Quantity:  product1.Quantity,
			CreatedAt: product1.CreatedAt,
		},
		{
			Id:        product2.Id,
			Name:      product2.Name,
			Price:     product2.Price,
			Quantity:  product2.Quantity,
			CreatedAt: product2.CreatedAt,
		},
	}

	assert.Equal(t, products, productResponses)

}
