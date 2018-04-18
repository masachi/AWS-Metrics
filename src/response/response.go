package response


type Response struct {
	MetricName string `json: "MetricName"`
	DateRegion []MetricDate `json: "dateRegion"`
	MetricData []float64 `json: "MetricData"`
}

type MetricDate struct {
	Timestamp interface{} `json: "timestamp"`
}