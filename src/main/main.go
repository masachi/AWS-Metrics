package main

import (
	"AWS-Metrics/src/router"
	"AWS-Metrics/src/instance"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
)

func main() {
	e := instance.GetInstance()
	router.Router()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
	}))
	e.Logger.Fatal(e.Start(":1323"))
}
