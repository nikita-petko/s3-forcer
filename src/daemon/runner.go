package daemon

import (
	"context"
	"sync"

	"github.com/golang/glog"
	"github.com/nikita-petko/s3-forcer/flags"
	"golang.org/x/sync/semaphore"
)

var runOnce sync.Once

// Run runs the doWork func in a loop as well
// as initializes the context.
// Can only occur once.
func Run() {
	runOnce.Do(func() {
		ctx := context.Background()
		sem := semaphore.NewWeighted(*flags.WorkerThreads)
		charlen := uint64(len(*flags.CharCombinations))
		length := *flags.MinLength
		chars := getChars()

		glog.Infof("Start work with %d semaphore threads.", *flags.WorkerThreads)

		for {
			doWork(ctx, sem, &charlen, &length, chars)
		}
	})
}
