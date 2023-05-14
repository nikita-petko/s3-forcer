package cache

import (
	"encoding/json"
	"io/fs"
	"os"
	"sync"

	"github.com/golang/glog"
	"github.com/nikita-petko/s3-forcer/flags"
)

var (
	setupOnce sync.Once
	positions map[string]int = map[string]int{}
	channels  []string       = []string{}

	defaultData *cachedData = &cachedData{Positions: map[string]int{}, Channels: []string{}}
)

// SetupCache sets up the caching environment
// by importing existing caches. Can only occur
// once.
func SetupCache() {
	setupOnce.Do(func() {
		var (
			fi  fs.FileInfo
			err error
		)

		fileName := computeCacheFileName()
		if *flags.InvalidateCache {
			glog.Warning("Trying to invalidate cache...")

			// It may or may not exist
			os.Remove(fileName)
		}

		if fi, err = os.Stat(fileName); os.IsNotExist(err) {
			t, _ := os.Create(fileName)

			serialized, err := json.Marshal(defaultData)
			if err != nil {
				glog.Fatal(err)
			}

			t.Write(serialized)

			t.Close()

			fi, _ = os.Stat(fileName)
		}

		if fi.IsDir() {
			glog.Fatalf("Cannot continue, %s is a directory", fileName)
		}

		contents, err := os.ReadFile(fi.Name())
		if err != nil {
			glog.Fatal(err)
		}

		var data cachedData
		err = json.Unmarshal(contents, &data)
		if err != nil {
			glog.Fatal(err)
		}

		positions = data.Positions
		channels = data.Channels
	})
}
