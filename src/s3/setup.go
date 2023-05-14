package s3

import (
	"net/http"
	"sync"
	"time"
)

var (
	s3HttpClient *http.Client = nil
	setupOnce    sync.Once
)

// SetupS3Client sets up the S3 HTTP client.
// Can only occur once.
func SetupS3Client() {
	setupOnce.Do(func() {
		s3HttpClient = http.DefaultClient
		s3HttpClient.Timeout = 15 * time.Second
	})
}
