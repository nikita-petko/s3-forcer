package s3

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/nikita-petko/s3-forcer/alerting"
	"github.com/nikita-petko/s3-forcer/cache"
	"github.com/nikita-petko/s3-forcer/flags"
	"github.com/nikita-petko/s3-forcer/metrics"
)

// ReadFromS3 reads from S3 using the channel
// version file via a HTTP HEAD to reduce overhead.
func ReadFromS3(channelCombination string) {
	channel := *flags.ChannelPrefix + channelCombination
	if cache.ChannelExists(channel) {
		glog.Warningf("Not sending redundant S3 HEAD when channel %s is already cached!", channel)
		return
	}

	format := rbxCdnProviderPerChannelVersion
	if *flags.UseS3Directly {
		format = s3PerChannelVersion
	}

	url := fmt.Sprintf(format, channel)

	resp, err := s3HttpClient.Head(url)
	if err != nil {
		glog.Fatalf("Error when sending request to S3: %v.", err)
	}

	if resp.StatusCode == 200 {
		glog.Infof("Found channel: %s", channel)

		metrics.NumberOfNewChannels.Inc()

		cache.RegisterChannel(channel)
		cache.FlushCache()

		go alerting.Alert(fmt.Sprintf("Found channel: %s", channel), fmt.Sprintf("Version Info: %s", url))
	}
}
