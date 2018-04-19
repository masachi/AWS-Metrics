package router

import (
	"AWS-Metrics/src/instance"
	"AWS-Metrics/src/controller"
)

func Router()  {
	e := instance.GetInstance()
	e.POST("aws/network/in", controller.NetWorkInController)
	e.OPTIONS("aws/network/in", controller.NetWorkInController)

	e.POST("aws/network/out", controller.NetworkOutController)
	e.OPTIONS("aws/network/out", controller.NetworkOutController)

	e.POST("aws/network/usage", controller.NetworkUsageController)
	e.OPTIONS("aws/network/usage", controller.NetworkUsageController)
}
