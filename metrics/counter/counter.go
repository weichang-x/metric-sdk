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
	counter := prometheus.NewCounterFrom(prom.CounterOpts{
		Namespace: nameSpace,
		Subsystem: subSystem,
		Name:      name,
		Help:      help,
	}, labels)
	return clientCounter{
		Name:    name,
		Counter: counter,
	}
}

func (client clientCounter) GetCounter() metrics.Counter {
	return client.Counter
}

func (client clientCounter) MetricName() string {
	return client.Name
}
