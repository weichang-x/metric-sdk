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

func NewGauge(nameSpace, subSystem, name string, help string, labels []string) Client {
	guage := prometheus.NewGaugeFrom(prom.GaugeOpts{
		Namespace: nameSpace,
		Subsystem: subSystem,
		Name:      name,
		Help:      help,
	}, labels)
	return clientGuage{
		Name:  name,
		Gauge: guage,
	}
}

func (client clientGuage) GetGuage() metrics.Gauge {
	return client.Gauge
}
func (client clientGuage) MetricName() string {
	return client.Name
}
