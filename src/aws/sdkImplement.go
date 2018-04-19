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
		Credentials: credentials.NewStaticCredentials("AKIAJMH6UZDI5UDV3PPQ", "5Hft/sR7rSIXz1QE85PPlD2QJ6KzYR0b74NgGyXS", ""),
	},
}))

var lightsailSvc = lightsail.New(awsSession)

func SdkImplement(param map[string]string, requestType string) lightsail.GetInstanceMetricDataOutput {
	switch requestType {
	case "in":
		return network(param, requestType)
		break
	case "out":
		return network(param, requestType)
		break
	}

	return lightsail.GetInstanceMetricDataOutput{}
}

func network(param map[string]string, requestType string) lightsail.GetInstanceMetricDataOutput {
	metricInput := lightsail.GetInstanceMetricDataInput{}
	startTime, _ := time.Parse("2006-01-02 15:04:05", param["startTime"])
	endTime, _ := time.Parse("2006-01-02 15:04:05", param["endTime"])

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

	return *metricData
}
