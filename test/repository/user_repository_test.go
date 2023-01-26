package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/faridlan/product-api/helper"
	"github.com/faridlan/product-api/model/domain"
	"github.com/faridlan/product-api/repository"
	"github.com/stretchr/testify/assert"
)

func getTx(db *sql.DB) *sql.Tx {
	tx, err := db.Begin()
	helper.PanicIfErr(err)

	return tx
}

func sqlMock() (*sql.DB, sqlmock.Sqlmock, error) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		return nil, nil, err
	} else {
		return db, mock, nil
	}

}

var productRepo = repository.NewProductRepository()

func TestProductRepoSuccess(t *testing.T) {
	db, mock, err := sqlMock()
	helper.PanicIfErr(err)
	defer db.Close()

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
	tx := getTx(db)
	mock.ExpectQuery("SELECT id, name, price, quantity, created_at, updated_at FROM product").WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"name",
		"price",
		"quantity",
		"created_at",
		"updated_at",
	}).AddRow(product1.Id, product1.Name, product1.Price, product1.Quantity, product1.CreatedAt, product1.UpdatedAt).AddRow(product2.Id, product2.Name, product2.Price, product2.Quantity, product2.CreatedAt, product2.UpdatedAt))
	mock.ExpectCommit()

	result := productRepo.FindAll(context.Background(), tx)

	products := []domain.Product{
		product1,
		product2,
	}

	assert.Equal(t, products, result)
}

func TestFindByIdSuccess(t *testing.T) {
	db, mock, err := sqlMock()
	helper.PanicIfErr(err)
	defer db.Close()

	product1 := domain.Product{
		Id:        1,
		Name:      "Laptop Asus",
		Price:     6000000,
		Quantity:  10,
		CreatedAt: time.Now().UnixMilli(),
		// UpdatedAt: &mysql.NullInt{},
	}

	mock.ExpectBegin()
	tx := getTx(db)
	mock.ExpectQuery("SELECT id, name, price, quantity, created_At, updated_at FROM product WHERE id = 1").WithArgs(1).WillReturnRows(sqlmock.NewRows(
		[]string{
			"id",
			"name",
			"price",
			"quantity",
			"created_at",
			"updated_at",
		},
	).AddRow(product1.Id, product1.Name, product1.Price, product1.Quantity, product1.CreatedAt, product1.UpdatedAt)),
}
