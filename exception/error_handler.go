package exception

import (
	"net/http"

	"github.com/faridlan/product-api/helper"
	"github.com/faridlan/product-api/model/web"
)

func ExceptionError(writer http.ResponseWriter, request *http.Request, err any) {

	if notFoundError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)

}

func internalServerError(writer http.ResponseWriter, request *http.Request, err any) {

	writer.Header().Add("content-type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func notFoundError(writer http.ResponseWriter, request *http.Request, err any) bool {

	exception, ok := err.(NotFoundError)

	if ok {
		writer.Header().Add("content-type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)

		webResponse := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}
