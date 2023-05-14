package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang/glog"
	"github.com/nikita-petko/s3-forcer/alerting"
	"github.com/nikita-petko/s3-forcer/cache"
	"github.com/nikita-petko/s3-forcer/daemon"
	"github.com/nikita-petko/s3-forcer/flags"
	"github.com/nikita-petko/s3-forcer/metrics"
	"github.com/nikita-petko/s3-forcer/s3"
)

var applicationName string
var buildMode string
var commitSha string

// Pre-setup, runs before main.
func init() {
	flags.SetupFlags(applicationName, buildMode, commitSha)
}

// Main entrypoint.
func main() {
	defer glog.Flush()

	if *flags.HelpFlag {
		flag.Usage()

		return
	}

	cache.SetupCache()
	alerting.SetupAlerting()
	s3.SetupS3Client()

	go metrics.SetupMetricsServer()
	go daemon.Run()

	// Wait for a signal to quit
	s := make(chan os.Signal, 1)

	// We want to catch ALL signals to quit
	signal.Notify(s, syscall.SIGABRT, syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		sig := <-s

		cache.FlushCache()

		glog.Warningf("Received signal %s, exiting\n", sig)

		os.Exit(0)
	}()
}
