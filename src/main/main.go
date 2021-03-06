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
		AllowHeaders: []string{"Origin", "No-Cache", "X-Requested-With", "If-Modified-Since", "Pragma", "Last-Modified", "Cache-Control", "Expires", "Content-Type", "X-E4M-With"},
	}))
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":1323"))
}
