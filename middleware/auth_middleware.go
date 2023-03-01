package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/faridlan/product-api/helper"
	"github.com/faridlan/product-api/helper/logging"
	"github.com/faridlan/product-api/model/web"
	"github.com/golang-jwt/jwt"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (authMiddleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	authorizationHeader := request.Header.Get("Authorization")

	// enpoints := exception.EndpointsGlobal()

	// for _, Enpoint := range enpoints {
	// 	if request.URL.Path != Enpoint.Url && request.Method != Enpoint.Method {
	// 		authMiddleware.Handler.ServeHTTP(writer, request)
	// 		return
	// 	}
	// }

	if request.URL.Path != "/api/seeder/products" && request.URL.Path != "/api/log" {
		authMiddleware.Handler.ServeHTTP(writer, request)
		return
	}

	if !strings.Contains(authorizationHeader, "Bearer") {

		writer.Header().Add("Content-Type", "application.json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		logging.ProductLogger(webResponse, writer, request)
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	var claim = &web.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("signing method invalid")
		}
		return web.JwtSecret, nil
	})

	if err != nil {
		writer.Header().Add("Content-Type", "application.json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		logging.ProductLogger(webResponse, writer, request)
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	if !token.Valid {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
		}

		helper.WriteToResponseBody(writer, webResponse)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	authMiddleware.Handler.ServeHTTP(writer, request)

}
