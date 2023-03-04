package main

import (
	"net/http"

	"github.com/faridlan/product-api/app"
	"github.com/faridlan/product-api/controller"
	"github.com/faridlan/product-api/exception"
	"github.com/faridlan/product-api/helper"
	"github.com/faridlan/product-api/middleware"
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

	router.POST("/api/products", ProductController.Create)
	router.POST("/api/seeder/products", ProductController.Seeder)
	router.PUT("/api/products/:productId", ProductController.Update)
	router.DELETE("/api/products/:productId", ProductController.Delete)
	router.DELETE("/api/seeder/products", ProductController.SeederDelete)
	router.GET("/api/products/:productId", ProductController.FindById)
	router.GET("/api/products", ProductController.FindAll)
	// router.GET("/api/log", ProductController.Logger)

	UserRepository := repository.NewUserRepository()

	UserService := service.NewAuthService(UserRepository, db)
	UserController := controller.NewUserControllerImpl(UserService)

	router.POST("/api/login", UserController.Login)

	router.PanicHandler = exception.ExceptionError

	server := http.Server{
		Addr:    ":9000",
		Handler: middleware.NewAuthMiddleware(router),
		// Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfErr(err)
}
