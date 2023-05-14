package flags

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/golang/glog"
)

var setupOnce sync.Once

// SetupFlags sets up flags for the current environment.
func SetupFlags(applicationName, buildMode, commitSha string) {
	setupOnce.Do(func() {
		flag.Usage = func() {
			os.Stderr.WriteString(fmt.Sprintf("Usage: %s\nBuild Mode: %s\nCommit: %s %s\n\n", applicationName, buildMode, commitSha, FlagsUsageString))
			flag.PrintDefaults()
		}

		flag.Set("logtostderr", "true")
		flag.Set("v", "100")

		flag.Parse()

		applyEnvironmentVariableFlags()

		if !*HelpFlag {
			glog.Infof("Flags setup complete!")
		}
	})
}
