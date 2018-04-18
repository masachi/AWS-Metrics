package router

import (
	"AWS-Metrics/src/instance"
	"AWS-Metrics/src/controller"
)

func Router()  {
	e := instance.GetInstance()
	e.POST("aws/network-in", controller.NetWorkInController)
	e.POST("aws/network-out", controller.NetworkOutController)
	e.get("aws/network-in", controller.NetworkOutController)
}
