package s3

import (
	"net/http"
	"sync"

	"github.com/nikita-petko/s3-forcer/flags"
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
		s3HttpClient.Timeout = *flags.S3ClientTimeout
	})
}
