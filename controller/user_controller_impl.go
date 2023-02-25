package controller

import (
	"net/http"

	"github.com/faridlan/product-api/helper"
	"github.com/faridlan/product-api/model/web"
	"github.com/faridlan/product-api/service"
	"github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.AuthService
}

func NewUserControllerImpl(userService service.AuthService) UserController {

	return &UserControllerImpl{
		UserService: userService,
	}

}

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {

	loginRequest := web.UserCreate{}
	helper.ReadFromRequestBody(request, &loginRequest)

	user, claim := controller.UserService.Login(request.Context(), loginRequest)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(web.JwtSecret)
	helper.PanicIfErr(err)

	user.Token = tokenString

	webRespone := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   user,
	}

	helper.WriteToResponseBody(writer, webRespone)
}
