package services

import (
	"AWS-Metrics/src/response"
	"reflect"
)

func NetWorkService(originData interface{}) interface{} {

	metricData := []float64{}
	dateRegion := []response.MetricDate{}

	networkData := reflect.ValueOf(originData).Elem().FieldByName("MetricData")

	for i := 0; i< networkData.Len(); i++ {
		date := response.MetricDate{}
		metricData = append(metricData, networkData.Index(i).Elem().FieldByName("Sum").Elem().Float() / 1024 / 1024 / 1024)
		date.Timestamp = networkData.Index(i).Elem().FieldByName("Timestamp").Elem().Interface()
		//date.Year = networkData.Index(i).Elem().FieldByName("Timestamp").Elem().String()[0:4]
		//date.Date = strings.Replace(networkData.Index(i).Elem().FieldByName("Timestamp").Elem().String()[5:10], "-", ".", -1)
		dateRegion = append(dateRegion, date)
	}

	netWorkResponse := response.Response{
		MetricName: reflect.ValueOf(originData).Elem().FieldByName("MetricName").Elem().String(),
		MetricData: metricData,
		DateRegion: dateRegion,
	}

	return netWorkResponse
}
