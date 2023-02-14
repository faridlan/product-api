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

	if validationError(writer, request, err) {
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

	exception, ok := err.(InterfaceError)

	if ok {
		writer.Header().Add("content-type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func validationError(writer http.ResponseWriter, request *http.Request, err any) bool {
	exception, ok := err.(ValidationError)
	// helper.TranslateError(err, helper.TranslateValidationEnglish())

	if ok {
		writer.Header().Add("content-type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		// english := english.New()
		// uni := ut.New(english, english)
		// trans, _ := uni.GetTranslator("en")

		// var errs []string
		// for _, err := range exception {
		// 	translatedError := err.Translate(trans)
		// 	errs = append(errs, translatedError)
		// }

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}
