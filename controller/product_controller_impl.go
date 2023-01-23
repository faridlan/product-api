package controller

import (
	"net/http"
	"strconv"

	"github.com/faridlan/product-api/helper"
	"github.com/faridlan/product-api/model/web"
	"github.com/faridlan/product-api/service"
	"github.com/julienschmidt/httprouter"
)

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

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	products := controller.ProductService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   products,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
