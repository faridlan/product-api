package controller

import (
	"embed"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/faridlan/product-api/helper"
	"github.com/faridlan/product-api/helper/logging"
	"github.com/faridlan/product-api/model/web"
	"github.com/faridlan/product-api/service"
	"github.com/julienschmidt/httprouter"
)

//go:embed json/products.json

var Json embed.FS

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	productCreate := web.ProductCreate{}
	helper.ReadFromRequestBody(request, &productCreate)

	product := controller.ProductService.Create(request.Context(), productCreate)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   product,
	}

	logging.ProductLogger(webResponse, writer, request)
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	productUpdate := web.ProductUpdate{}
	helper.ReadFromRequestBody(request, &productUpdate)

	idProduct := param.ByName("id")
	id, err := strconv.Atoi(idProduct)
	helper.PanicIfErr(err)

	productUpdate.Id = id

	product := controller.ProductService.Update(request.Context(), productUpdate)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   product,
	}

	logging.ProductLogger(webResponse, writer, request)
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	idProduct := param.ByName("id")
	id, err := strconv.Atoi(idProduct)
	helper.PanicIfErr(err)

	controller.ProductService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	logging.ProductLogger(webResponse, writer, request)
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	idProduct := param.ByName("id")
	id, err := strconv.Atoi(idProduct)
	helper.PanicIfErr(err)

	product := controller.ProductService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   product,
	}

	logging.ProductLogger(webResponse, writer, request)
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	products := controller.ProductService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   products,
	}

	logging.ProductLogger(webResponse, writer, request)
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) Seeder(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {

	products, err := Json.ReadFile("json/products.json")
	helper.PanicIfErr(err)

	productCreate := []web.ProductCreate{}
	json.Unmarshal(products, &productCreate)

	product := controller.ProductService.CreateMany(request.Context(), productCreate)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   product,
	}

	logging.ProductLogger(webResponse, writer, request)
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) SeederDelete(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	controller.ProductService.DeleteAll(request.Context())

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	logging.ProductLogger(webResponse, writer, request)

	helper.WriteToResponseBody(writer, webResponse)
}
