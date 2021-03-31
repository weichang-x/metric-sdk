package counter

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
	prom "github.com/prometheus/client_golang/prometheus"
)

type clientCounter struct {
	metrics.Counter
	Name string
}

func NewCounter(nameSpace, subSystem, name string, help string, conLabels map[string]string) Client {
	var labels []string
	var labelValues []string
	if len(conLabels) > 0 {
		for label, value := range conLabels {
			labels = append(labels, label)
			labelValues = append(labelValues, label, value)
		}
	}
	counter := prometheus.NewCounterFrom(
		prom.CounterOpts{
			Namespace: nameSpace,
			Subsystem: subSystem,
			Name:      name,
			Help:      help,
		},
		labels)
	if len(labels) > 0 {
		counter.With(labelValues...)
	}
	return clientCounter{
		Name:    name,
		Counter: counter,
	}
}

func (client clientCounter) Add(value float64) {
	client.Counter.Add(value)
}

func (client clientCounter) MetricName() string {
	return client.Name
}
