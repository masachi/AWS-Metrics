package main

import (
	"AWS-Metrics/src/router"
	"AWS-Metrics/src/instance"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := instance.GetInstance()
	router.Router()
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	e.Logger.Fatal(e.Start(":1323"))
}
