package response

import "go/types"

type Response struct {
	metricName string `json: "MetricName"`
	dateRegion types.Array `json: "dateRegion"`
	metricData types.Array `json: "MetricData"`
}

