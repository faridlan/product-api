package main

import (
	"net/http"

	"github.com/faridlan/product-api/app"
	"github.com/faridlan/product-api/controller"
	"github.com/faridlan/product-api/exception"
	"github.com/faridlan/product-api/helper"
	"github.com/faridlan/product-api/repository"
	"github.com/faridlan/product-api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()
	validate := validator.New()
	db := app.NewDatabase()

	ProductRepository := repository.NewProductRepository()
	ProductService := service.NewProductService(ProductRepository, db, validate)
	ProductController := controller.NewProductController(ProductService)

	router.POST("/api/products/", ProductController.Create)
	router.PUT("/api/products/:id", ProductController.Update)
	router.DELETE("/api/products/:id", ProductController.Delete)
	router.GET("/api/products/:id", ProductController.FindById)
	router.GET("/api/products/", ProductController.FindAll)

	router.PanicHandler = exception.ExceptionError

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfErr(err)
}
