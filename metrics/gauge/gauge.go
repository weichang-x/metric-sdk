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
	guage := prometheus.NewGaugeFrom(prom.GaugeOpts{
		Namespace: nameSpace,
		Subsystem: subSystem,
		Name:      name,
		Help:      help,
	}, labels)

	if len(labels) > 0 {
		guage.With(labelValues...)
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
