package gauge

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
	prom "github.com/prometheus/client_golang/prometheus"
)

type clientGuage struct {
	Name string
	metrics.Gauge
}

func NewGauge(nameSpace, subSystem, name string, help string, conLabels map[string]string) Client {
	var labels []string
	var labelValues []string
	if len(conLabels) > 0 {
		for label, value := range conLabels {
			labels = append(labels, label)
			labelValues = append(labelValues, label, value)
		}
	}

	var guage metrics.Gauge
	gaugeOpts := prom.GaugeOpts{
		Namespace: nameSpace,
		Subsystem: subSystem,
		Name:      name,
		Help:      help,
	}

	if len(labels) > 0 {
		guage = prometheus.NewGaugeFrom(gaugeOpts, labels).With(labelValues...)
	} else {
		guage = prometheus.NewGaugeFrom(gaugeOpts, labels)
	}
	return clientGuage{
		Name:  name,
		Gauge: guage,
	}
}

func (client clientGuage) Set(value float64) {
	client.Gauge.Set(value)
}

func (client clientGuage) MetricName() string {
	return client.Name
}
