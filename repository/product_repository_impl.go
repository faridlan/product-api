package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/product-api/helper"
	"github.com/faridlan/product-api/model/domain"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "INSERT INTO product(name, price, quantity, created_at) values (?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.Quantity, product.CreatedAt)
	helper.PanicIfErr(err)

	id, err := result.LastInsertId()
	helper.PanicIfErr(err)

	product.Id = int(id)

	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "UPDATE product SET name = ?, price = ?, quantity = ?, updated_at = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.Quantity, product.UpdatedAt, product.Id)
	helper.PanicIfErr(err)

	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	SQL := "DELETE FROM product WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.PanicIfErr(err)
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error) {
	SQL := "SELECT id, name, price, quantity, created_at, updated_at FROM product WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfErr(err)

	defer rows.Close()

	product := domain.Product{}

	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.CreatedAt, &product.UpdatedAt)
		helper.PanicIfErr(err)

		return product, nil
	} else {
		return product, errors.New("product not found")
	}
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	SQL := "SELECT id, name, price, quantity, created_at, updated_at FROM product"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfErr(err)

	defer rows.Close()
	products := []domain.Product{}

	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.CreatedAt, &product.UpdatedAt)
		helper.PanicIfErr(err)

		products = append(products, product)
	}

	return products
}
