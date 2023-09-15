package auth

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server-golang/models/request"
	"server-golang/models/response"
	s_user "server-golang/services"
)

func Register(context echo.Context) error {
	body := new(request.Register)
	context.Bind(body)

	save := s_user.SaveUser(body)
	return context.JSON(http.StatusCreated, save)
}
func Login(context echo.Context) error {
	return context.JSON(http.StatusCreated, response.Response{
		Status:  true,
		Message: "Login success",
		Data:    response.Token{},
	})
}

func Logout(context echo.Context) error {

	return context.JSON(http.StatusCreated, response.Response{
		Status:  true,
		Message: "Logout success",
		Data:    nil,
	})
}
