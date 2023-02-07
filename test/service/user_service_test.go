package service_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/faridlan/product-api/helper"
	"github.com/faridlan/product-api/helper/mysql"
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
	defer db.Close()

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

func TestFindbyIdSuccess(t *testing.T) {
	db, mock, err := sqlMock()
	helper.PanicIfErr(err)
	validate := validator.New()
	defer db.Close()

	ProductService := NewService(db, validate)

	product1 := domain.Product{
		Id:        1,
		Name:      "Laptop Asus",
		Price:     6000000,
		Quantity:  10,
		CreatedAt: time.Now().UnixMilli(),
		// UpdatedAt: &mysql.NullInt{},
	}

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT id, name, price, quantity, created_at, updated_at FROM product WHERE id = ?").WithArgs(1).WillReturnRows(sqlmock.NewRows(
		[]string{
			"id",
			"name",
			"price",
			"quantity",
			"created_at",
			"updated_at",
		},
	).AddRow(product1.Id, product1.Name, product1.Price, product1.Quantity, product1.CreatedAt, product1.UpdatedAt))
	mock.ExpectCommit()

	product := ProductService.FindById(context.Background(), product1.Id)

	productResponse := web.ProductResponse{
		Id:        product1.Id,
		Name:      product1.Name,
		Price:     product1.Price,
		Quantity:  product1.Quantity,
		CreatedAt: product1.CreatedAt,
	}

	assert.Equal(t, product, productResponse)
}

func TestFindByIdFailed(t *testing.T) {

	db, mock, err := sqlMock()
	helper.PanicIfErr(err)
	validate := validator.New()
	defer db.Close()

	ProductService := NewService(db, validate)

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT id, name, price, quantity, created_at, updated_at FROM product WHERE id = ?").WithArgs(2).WillReturnRows(sqlmock.NewRows(
		[]string{
			"id",
			"name",
			"price",
			"quantity",
			"created_at",
			"updated_at",
		},
	))
	mock.ExpectRollback()

	product := ProductService.FindById(context.Background(), 2)

	assert.Empty(t, product)
}

func TestCreate(t *testing.T) {

	db, mock, err := sqlMock()
	helper.PanicIfErr(err)
	validate := validator.New()
	defer db.Close()

	ProductService := NewService(db, validate)

	product := web.ProductCreate{
		Name:      "Laptop Lenovo",
		Price:     7000000,
		Quantity:  10,
		CreatedAt: time.Now().UnixMilli(),
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO product(name, price, quantity, created_at) values (?,?,?,?)").WithArgs(product.Name, product.Price, product.Quantity, product.CreatedAt).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	productResponse := ProductService.Create(context.Background(), product)

	productExpec := web.ProductResponse{
		Id:        1,
		Name:      product.Name,
		Price:     product.Price,
		Quantity:  product.Quantity,
		CreatedAt: product.CreatedAt,
	}

	assert.Equal(t, productExpec, productResponse)

}

func TestUpdate(t *testing.T) {

	db, mock, err := sqlMock()
	helper.PanicIfErr(err)
	validate := validator.New()

	defer db.Close()

	ProductService := NewService(db, validate)

	productUpdate := web.ProductUpdate{
		Id:        1,
		Name:      "Laptop Lenovo Core i3",
		Price:     8000000,
		Quantity:  15,
		UpdatedAt: time.Now().UnixMilli(),
	}

	product1 := domain.Product{
		Id:        1,
		Name:      "Laptop Asus",
		Price:     6000000,
		Quantity:  10,
		CreatedAt: time.Now().UnixMilli(),
		// UpdatedAt: &mysql.NullInt{},
	}

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT id, name, price, quantity, created_at, updated_at FROM product WHERE id = ?").WithArgs(1).WillReturnRows(sqlmock.NewRows(
		[]string{
			"id",
			"name",
			"price",
			"quantity",
			"created_at",
			"updated_at",
		},
	).AddRow(product1.Id, product1.Name, product1.Price, product1.Quantity, product1.CreatedAt, product1.UpdatedAt))
	mock.ExpectExec("UPDATE product SET name = ?, price = ?, quantity = ?, updated_at = ? where id = ?").WithArgs(productUpdate.Name, productUpdate.Price, productUpdate.Quantity, productUpdate.UpdatedAt, productUpdate.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	product := ProductService.Update(context.Background(), productUpdate)

	productExpec := web.ProductResponse{
		Id:        1,
		Name:      productUpdate.Name,
		Price:     productUpdate.Price,
		Quantity:  productUpdate.Quantity,
		CreatedAt: product1.CreatedAt,
		UpdatedAt: &mysql.NullInt{
			NullInt64: sql.NullInt64{
				Int64: productUpdate.UpdatedAt,
				Valid: true,
			},
		},
	}

	assert.Equal(t, productExpec, product)

}

func TestDelete(t *testing.T) {

	db, mock, err := sqlMock()
	helper.PanicIfErr(err)
	validate := validator.New()
	defer db.Close()

	ProductService := NewService(db, validate)

	product := domain.Product{
		Id: 1,
	}

	product1 := domain.Product{
		Id:        1,
		Name:      "Laptop Asus",
		Price:     6000000,
		Quantity:  10,
		CreatedAt: time.Now().UnixMilli(),
		// UpdatedAt: &mysql.NullInt{},
	}

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT id, name, price, quantity, created_at, updated_at FROM product WHERE id = ?").WithArgs(1).WillReturnRows(sqlmock.NewRows(
		[]string{
			"id",
			"name",
			"price",
			"quantity",
			"created_at",
			"updated_at",
		},
	).AddRow(product1.Id, product1.Name, product1.Price, product1.Quantity, product1.CreatedAt, product1.UpdatedAt))
	mock.ExpectExec("DELETE FROM product WHERE id = ?").WithArgs(product.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	ProductService.Delete(context.Background(), product.Id)
}
