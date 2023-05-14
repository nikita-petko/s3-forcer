package sns

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	awssns "github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/golang/glog"
)

var (
	snsClient *awssns.Client = nil
	setupOnce sync.Once
)

// SetupSnsTopic sets up the SNS topic.
// Can only occur once.
func SetupSnsTopic() {
	setupOnce.Do(func() {
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			glog.Fatalf("Error initializing AWS SNS Config: %v", err)
		}

		snsClient = awssns.NewFromConfig(cfg)

		glog.Info("AWS SNS setup complete!")
	})
}
