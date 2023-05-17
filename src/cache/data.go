package cache

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/golang/glog"
	"github.com/nikita-petko/s3-forcer/flags"
	"github.com/nikita-petko/s3-forcer/metrics"
)

// FlushCache flushes all memory cached values to the cache file.
func FlushCache() {
	if !*flags.UseCache {
		return
	}

	file, _ := os.OpenFile(computeCacheFileName(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer file.Close()

	file.Truncate(0)
	file.Seek(0, 0)

	w := bufio.NewWriter(file)

	data := &cachedData{
		Positions: *positions,
		Channels:  *channels,
	}

	b, err := json.Marshal(data)
	if err != nil {
		glog.Fatal(err)
	}

	w.Write(b)

	w.Flush()
}

// RegisterChannel registers a channel into the cached channels.
func RegisterChannel(channel string) {
	metrics.RegisteredChannels.WithLabelValues(channel).Set(1)

	glog.Infof("Registered new channel %s", channel)

	if *flags.UseCache {
		*channels = append(*channels, channel)
	}
}

// PositionExists checks if the specified postion exists.
func PositionExists(length string) (exists bool, pos uint64) {
	if !*flags.UseCache {
		return false, 0
	}

	if len(*positions) == 0 {
		return false, 0
	}

	for k, v := range *positions {
		if k == length {
			return true, v
		}
	}

	return false, 0
}

// SetCachedPosition sets the cached lstr value.
func SetCachedPostion(lstr string, position uint64) {
	if *flags.UseCache {
		(*positions)[lstr] = position
	}
}

// ChannelExists checks if the specified channel exists in the cache.
func ChannelExists(channel string) (exists bool) {
	if !*flags.UseCache {
		return false
	}

	if len(*channels) == 0 {
		return false
	}

	for _, v := range *channels {
		if v == channel {
			return true
		}
	}

	return false
}
