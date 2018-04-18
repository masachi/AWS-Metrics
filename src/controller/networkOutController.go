package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"AWS-Metrics/src/aws"
)

func NetworkOutController(context echo.Context) error {
	//startTime := context.FormValue("startTime")
	//endTime := context.FormValue("endTime")

	params := make(map[string]string)
	//params["startTime"] = startTime
	//params["endTime"] = endTime
	params["startTime"] = "2018-03-27 00:00:00"
	params["endTime"] = "2018-04-18 00:00:00"


	return context.JSON(http.StatusOK, aws.SdkImplement(params, "out"))
}
