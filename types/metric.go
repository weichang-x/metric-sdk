package types

import (
	"time"
)

type (
	ScrapMetric map[string]interface{}
)

type Metrics struct {
	Timestamp time.Time     `json:"time_stamp"`
	Type      string        `json:"type"`
	Help      string        `json:"help,omitempty"`
	Metrics   []ScrapMetric `json:"metrics"`
}

// MetricGroup adds the grouping labels to a NameToTimestampedMetricFamilyMap.
type MetricGroup struct {
	Labels  map[string]string
	Metrics []*GobbableMetricFamily
}
