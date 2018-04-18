package aws

import (
	"github.com/aws/aws-sdk-go/service/lightsail"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"time"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

var awsSession = session.Must(session.NewSessionWithOptions(session.Options{
	Config: aws.Config{
		Region: aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials("AKIAIJZD3NLWWZ4D4ZHA", "jjTsMUb4XcCHFL+0WbBreiTpXF2/p74vLnVI2MZJ", ""),
	},
}))

var lightsailSvc = lightsail.New(awsSession)

func SdkImplement(param map[string]string, requestType string) interface{} {
	switch requestType {
	case "in":
		return network(param, requestType)
		break
	case "out":
		return network(param, requestType)
		break
	}

	return []byte{}
}

func network(param map[string]string, requestType string) interface{} {
	metricInput := lightsail.GetInstanceMetricDataInput{}
	startTime, _ := time.Parse("2006-01-02 15:04:05", "2018-03-27 00:00:00")
	endTime := time.Now()

	metricInput.SetInstanceName("Debian-512MB-Tokyo-1")
	switch requestType {
	case "in":
		metricInput.SetMetricName("NetworkIn")
		break
	case "out":
		metricInput.SetMetricName("NetworkOut")
		break
	}
	metricInput.SetPeriod(86400)
	metricInput.SetStatistics([]*string{
		aws.String("Sum"),
	})
	metricInput.SetStartTime(startTime)
	metricInput.SetEndTime(endTime)
	metricInput.SetUnit("Bytes")

	metricData, _  := lightsailSvc.GetInstanceMetricData(&metricInput)

	return metricData
}
