package response


type Network struct {
	MetricName string `json:"MetricName"`
	DateRegion []MetricDate `json:"dateRegion"`
	MetricData []float64 `json:"MetricData"`
}

type MetricDate struct {
	Year string `json:"year"`
	Date string `json:"date"`
}

type NetworkUsage struct {
	MetricName string `json:"MetricName"`
	MetricData Bandwidth  `json:"MetricData"`
}

type Bandwidth struct {
	All float64 `json:"All"`
	Out float64 `json:"Out"`
	In float64 `json:"In"`
}