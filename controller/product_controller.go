package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductController interface {
	Create(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Seeder(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	SeederDelete(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	// Logger(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}
