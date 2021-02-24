package types

import "github.com/prometheus/client_golang/prometheus"

type (
	Counter prometheus.Counter
	Gauge   prometheus.Gauge
)
