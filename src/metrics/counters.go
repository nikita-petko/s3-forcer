package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// The number of new channels.
var NumberOfNewChannels = promauto.NewCounter(
	prometheus.CounterOpts{
		Name: "number_of_new_channels",
		Help: "The number of new channels",
	},
)

// The number of existing channels.
var NumberOfExistingChannels = promauto.NewCounter(
	prometheus.CounterOpts{
		Name: "number_of_existing_channels",
		Help: "The number of existing channels",
	},
)

// The currently registered channels, either from the remote service or the cache.
var RegisteredChannels = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "registered_channels",
		Help: "The currently registered channels, either from the remote service or the cache.",
	},
	[]string{"channel_name"},
)
