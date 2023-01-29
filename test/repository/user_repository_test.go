package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/faridlan/product-api/helper"
	"github.com/faridlan/product-api/helper/mysql"
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

	result, err := productRepo.FindById(context.Background(), tx, 1)
	productRespone := domain.Product{
		Id:        1,
		Name:      "Laptop Asus",
		Price:     6000000,
		Quantity:  10,
		CreatedAt: time.Now().UnixMilli(),
	}

	assert.Nil(t, err)
	assert.Equal(t, result, productRespone)
}

func TestFindByIdFailed(t *testing.T) {
	db, mock, err := sqlMock()
	helper.PanicIfErr(err)
	defer db.Close()

	mock.ExpectBegin()
	tx := getTx(db)
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

	product, err := productRepo.FindById(context.Background(), tx, 2)

	assert.NotNil(t, err)
	assert.Empty(t, product)
	assert.Error(t, err)
}

func TestSave(t *testing.T) {
	db, mock, err := sqlMock()
	helper.PanicIfErr(err)

	defer db.Close()

	product := domain.Product{
		Name:      "Laptop Lenovo",
		Price:     7000000,
		Quantity:  10,
		CreatedAt: time.Now().UnixMilli(),
	}

	mock.ExpectBegin()
	tx := getTx(db)
	mock.ExpectExec("INSERT INTO product(name, price, quantity, created_at) values (?,?,?,?)").WithArgs(product.Name, product.Price, product.Quantity, product.CreatedAt).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	productExpec := domain.Product{
		Id:        1,
		Name:      "Laptop Lenovo",
		Price:     7000000,
		Quantity:  10,
		CreatedAt: time.Now().UnixMilli(),
	}

	productResult := productRepo.Save(context.Background(), tx, productExpec)

	assert.Equal(t, productResult, productExpec)
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlMock()
	helper.PanicIfErr(err)
	defer db.Close()

	productUpdate := domain.Product{
		Id:       1,
		Name:     "Laptop Lenovo Core i3",
		Price:    8000000,
		Quantity: 15,
		UpdatedAt: &mysql.NullInt{
			NullInt64: sql.NullInt64{
				Int64: 1675010292187,
				Valid: true,
			},
		},
	}

	mock.ExpectBegin()
	tx := getTx(db)
	mock.ExpectExec("UPDATE product SET name = ?, price = ?, quantity = ?, updated_at = ? where id = ?").WithArgs(productUpdate.Name, productUpdate.Price, productUpdate.Quantity, productUpdate.UpdatedAt, productUpdate.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	productExpec := domain.Product{
		Id:        1,
		Name:      "Laptop Lenovo Core i3",
		Price:     8000000,
		Quantity:  15,
		CreatedAt: 1675010173412,
		UpdatedAt: &mysql.NullInt{
			NullInt64: sql.NullInt64{
				Int64: 1675010292187,
				Valid: true,
			},
		},
	}

	product := productRepo.Update(context.Background(), tx, productExpec)

	assert.Equal(t, product, productExpec)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlMock()
	helper.PanicIfErr(err)
	defer db.Close()

	product := domain.Product{
		Id: 1,
	}

	mock.ExpectBegin()
	tx := getTx(db)
	mock.ExpectExec("DELETE FROM product WHERE id = ?").WithArgs(product.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	productRepo.Delete(context.Background(), tx, product)
}
