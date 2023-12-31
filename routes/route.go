package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	auth "server-golang/controllers/auth"
	user "server-golang/controllers/auth"
	mid "server-golang/middleware"
)

func Init(app *echo.Echo) {
	app.Use(middleware.Logger())
	// Tanpa validasi auth
	app.POST("/login", auth.Login)
	app.POST("/register", auth.Register)
	app.POST("/logout", auth.Logout)

	// Endpoint perlu validasi auth
	app.POST("/api/users/profile", mid.JwtAuth(user.AddProfile))
	app.PUT("/api/users/profile", mid.JwtAuth(user.UpdateProfile))

	app.POST("/api/users/address", mid.JwtAuth(user.AddAddress))
	app.PUT("/api/users/address", mid.JwtAuth(user.UpdateAddress))

	app.GET("/api/users", mid.JwtAuth(user.GetUser))

}
