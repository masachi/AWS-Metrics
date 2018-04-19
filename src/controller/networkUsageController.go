package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"AWS-Metrics/src/services"
)

func NetworkUsageController(context echo.Context) error {
	startTime := context.FormValue("startTime")
	endTime := context.FormValue("endTime")

	params := make(map[string]string)
	params["startTime"] = startTime
	params["endTime"] = endTime

	return context.JSON(http.StatusOK, services.NetworkUsageService(params))
}
