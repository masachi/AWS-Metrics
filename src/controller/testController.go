package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

func TestController(context echo.Context) error {
	return context.JSON(http.StatusOK, "2333333")
}
