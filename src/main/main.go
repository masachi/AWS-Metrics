package main

import (
	"AWS-Metrics/src/router"
	"AWS-Metrics/src/instance"
)

func main() {
	e := instance.GetInstance()
	router.Router()
	e.Logger.Fatal(e.Start(":1323"))
}
