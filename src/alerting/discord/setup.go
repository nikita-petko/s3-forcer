package discord

import (
	"net/http"
	"sync"
	"time"

	"github.com/golang/glog"
)

var (
	discordHttpClient *http.Client = nil
	setupOnce         sync.Once
)

// SetupDisciordClient sets up the Discord HTTP Client.
// Can only occur once.
func SetupDiscordClient() {
	setupOnce.Do(func() {
		discordHttpClient = http.DefaultClient
		discordHttpClient.Timeout = 15 * time.Second

		glog.Info("Discord WebHook setup complete!")
	})
}
