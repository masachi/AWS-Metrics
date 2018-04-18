package controller

import (
	"github.com/labstack/echo"
	"AWS-Metrics/src/aws"
	"net/http"
)

func NetWorkInController(context echo.Context) error {
	startTime := context.FormValue("startTime")
	endTime := context.FormValue("endTime")

	params := make(map[string]string)
	params["startTime"] = startTime
	params["endTime"] = endTime


	return context.JSON(http.StatusOK, aws.SdkImplement(params, "in"))
}
