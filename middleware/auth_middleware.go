package middleware

import (
	"net/http"
	"strings"

	"github.com/faridlan/product-api/helper"
	"github.com/faridlan/product-api/helper/logging"
	"github.com/faridlan/product-api/model/web"
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

	// Endpoints := exception.EndpointsAdmin()

	// for _, Enpoint := range Endpoints {
	if request.URL.Path == "/api/products" {
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

	// tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	// user := domain.User{
	// 	Username: "Admin",
	// 	Password: "admin",
	// }

	// var claim = &web.Claims{}

	// token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
	// 	if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf("signing method invalid")
	// 	} else if method != web.JwtSigningMEethod {
	// 		return nil, fmt.Errorf("signing method invalid")
	// 	}
	// 	return web.JwtSecret, nil
	// })

	// if err != nil {
	// 	writer.Header().Add("Content-Type", "application.json")
	// 	writer.WriteHeader(http.StatusUnauthorized)

	// 	webResponse := web.WebResponse{
	// 		Code:   http.StatusUnauthorized,
	// 		Status: "UNAUTHORIZED",
	// 	}

	// 	logging.ProductLogger(webResponse, writer, request)
	// 	helper.WriteToResponseBody(writer, webResponse)
	// 	return
	// }

}
