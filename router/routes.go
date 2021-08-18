package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"intra_payment/controllers"
)

func NewRouter() *echo.Echo{

	e:= echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	base := e.Group("/server/v1")

	base.GET("/", controllers.Home)

	return e
}
