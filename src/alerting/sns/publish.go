package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/golang/glog"
	"github.com/nikita-petko/s3-forcer/flags"
)

// PublishToSnsTopic publishes a message to the configured SNS topic.
func PublishToSnsTopic(body string) {
	if snsClient == nil {
		return
	}

	_, err := snsClient.Publish(context.TODO(), &sns.PublishInput{
		TopicArn: flags.SnsTopicArn,
		Message:  &body,
	})

	if err != nil {
		glog.Errorf("Error occurred when publishing to SNS Topic: %v (topic ARN: %s)", err, *flags.SnsTopicArn)
	}
}
