package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func Request ( c echo.Context ) error {








	return c.JSON(http.StatusOK, "requested")
}