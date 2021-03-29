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

func NewCounter(nameSpace, subSystem, name string, help string, labels []string) Client {
	return clientCounter{
		Name: name,
		Counter: prometheus.NewCounter(prom.NewCounterVec(
			prom.CounterOpts{
				Namespace: nameSpace,
				Subsystem: subSystem,
				Name:      name,
				Help:      help,
			},
			labels)),
	}
}

func (client clientCounter) Add(value float64) {
	client.Counter.Add(value)
}

func (client clientCounter) MetricName() string {
	return client.Name
}
