package metrics

import (
	"net/http"
	"sync"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/nikita-petko/s3-forcer/flags"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var setupOnce sync.Once

// SetupMetrics server sets up the prometheus metrics server.
// Can only occur once.
func SetupMetricsServer() {
	setupOnce.Do(func() {
		glog.Infof("Starting Prometheus metrics server with bind address: %s", *flags.BindAddressIpv4)

		router := mux.NewRouter()

		prometheus.Unregister(collectors.NewGoCollector())
		prometheus.Unregister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
		router.PathPrefix("/").HandlerFunc(promhttp.Handler().ServeHTTP)

		if err := http.ListenAndServe(*flags.BindAddressIpv4, router); err != nil {
			glog.Errorf("Error when starting Prometheus metrics server: %v", err)
		}
	})
}
