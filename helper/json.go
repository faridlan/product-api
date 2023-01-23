package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result any) {
	decode := json.NewDecoder(request.Body)
	err := decode.Decode(result)
	PanicIfErr(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response any) {
	writer.Header().Add("content-type", "application/json")
	encode := json.NewEncoder(writer)
	err := encode.Encode(response)
	PanicIfErr(err)
}
