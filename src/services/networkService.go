package services

import (
	"AWS-Metrics/src/response"
	"AWS-Metrics/src/aws"
	"github.com/aws/aws-sdk-go/service/lightsail"
	"strconv"
)

func NetWorkService(originData lightsail.GetInstanceMetricDataOutput) interface{} {

	metricData := []float64{}
	dateRegion := []response.MetricDate{}

	networkData := SortService(originData.MetricData)

	for _, item := range networkData {
		date := response.MetricDate{}
		metricData = append(metricData, *item.Sum / 1024 / 1024 / 1024)
		date.Year = strconv.Itoa((*item.Timestamp).Year())
		date.Date = (*item.Timestamp).String()[6:10]
		dateRegion = append(dateRegion, date)
	}

	netWorkResponse := response.Network{
		MetricName: *originData.MetricName,
		MetricData: metricData,
		DateRegion: dateRegion,
	}

	return netWorkResponse
}

func NetworkUsageService(params map[string]string) interface{} {
	in := countBandWidth(aws.SdkImplement(params, "in"))
	out := countBandWidth(aws.SdkImplement(params, "out"))

	return response.NetworkUsage{
		MetricName: "All",
		MetricData: response.Bandwidth{
			All: in + out,
			In: in,
			Out: out,
		},
	}
}

func countBandWidth(originData lightsail.GetInstanceMetricDataOutput) float64 {
	var usage float64
	usage = 0
	networkData := originData.MetricData
	for _, item := range networkData {
		usage += *item.Sum / 1024 / 1024 / 1024
	}

	return usage
}
