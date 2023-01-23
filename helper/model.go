package helper

import (
	"github.com/faridlan/product-api/model/domain"
	"github.com/faridlan/product-api/model/web"
)

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id:        product.Id,
		Name:      product.Name,
		Price:     product.Price,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}

func ToProductResponses(products []domain.Product) []web.ProductResponse {
	productResponse := []web.ProductResponse{}
	for _, product := range products {
		productResponse = append(productResponse, ToProductResponse(product))
	}

	return productResponse
}
