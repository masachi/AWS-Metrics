package services

import (
	"github.com/aws/aws-sdk-go/service/lightsail"
	"sort"
)

type DataPoints []lightsail.MetricDatapoint

func (points DataPoints) Len() int {
	return len(points)
}

func (points DataPoints) Less(i, j int) bool {
	return points[i].Timestamp.Before(*points[j].Timestamp)
}

func (points DataPoints) Swap(i, j int)  {
	points[i], points[j] = points[j], points[i]
}

func SortService(originData []*lightsail.MetricDatapoint) DataPoints {
	data := DataPoints{}
	for _, item := range originData {
		data = append(data, *item)
	}
	sort.Sort(data)
	return data
}
