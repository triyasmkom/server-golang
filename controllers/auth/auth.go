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
	save := s_user.Register(body)
	return context.JSON(http.StatusCreated, save)
}
func Login(context echo.Context) error {
	body := new(request.Login)
	context.Bind(body)
	save := s_user.Login(body)
	return context.JSON(http.StatusCreated, save)
}

// not use
func Logout(context echo.Context) error {

	return context.JSON(http.StatusCreated, response.Response{
		Status:  true,
		Message: "Logout success",
		Data:    response.Data{},
	})
}

func AddProfile(context echo.Context) error {
	addProfile := s_user.AddProfile(context)
	return context.JSON(http.StatusCreated, addProfile)
}

func UpdateProfile(context echo.Context) error {

	return context.JSON(http.StatusCreated, response.Response{
		Status:  true,
		Message: "Update success",
		Data:    response.Data{},
	})
}
