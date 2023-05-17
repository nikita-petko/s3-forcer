package daemon

import (
	"context"
	"math"
	"strconv"
	"time"

	"github.com/golang/glog"
	"github.com/nikita-petko/s3-forcer/cache"
	"github.com/nikita-petko/s3-forcer/flags"
	"github.com/nikita-petko/s3-forcer/itertools"
	"github.com/nikita-petko/s3-forcer/s3"
	"golang.org/x/sync/semaphore"
)

func doWork(ctx context.Context, sem *semaphore.Weighted, charlen, length *int, chars []interface{}) {
	glog.Infof("Starting attempts @ %d chars (max attempts: %d)", *length, int64(math.Pow(float64(*charlen), float64(*length))))
	t := time.Now()

	channelNames := joinTogether(itertools.Product(*length, chars))

	currentAttempt := 1

	lenstr := strconv.Itoa(*length)

	if exists, pos := cache.PositionExists(lenstr); exists {
		glog.Infof("Starting at cached position of %d", pos)

		currentAttempt = pos
	}

	for _, channel := range channelNames {
		if currentAttempt%*flags.AttemptMilestone == 0 {
			glog.Infof("Milestone of %d attempts reached.", currentAttempt)

			cache.SetCachedPostion(lenstr, currentAttempt)
			cache.FlushCache()
		}

		if err := sem.Acquire(ctx, 1); err != nil {
			glog.Fatalf("Error acquiring semaphore lock: %v", err)
		}

		go func(channel string) {
			s3.ReadFromS3(channel)
			sem.Release(1)
		}(channel)

		currentAttempt++
	}

	glog.Infof("Cycle for @ %d chars took: %s", *length, time.Since(t).String())

	*length++
}
