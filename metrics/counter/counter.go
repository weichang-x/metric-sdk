package counter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/weichang-bianjie/metric-sdk/core"
	"github.com/weichang-bianjie/metric-sdk/types"
)

func NewCounter() Client {
	return clientCounter{}
}

type clientCounter struct {
	prometheus.Counter
}

func (client clientCounter) RegisterMetricInfo(name string, help string, constLabels map[string]string) types.Counter {
	completionOpts := prometheus.NewCounter(prometheus.CounterOpts{
		Name:        name,
		Help:        help,
		ConstLabels: constLabels,
	})
	core.RegisterCollector(completionOpts)
	return completionOpts
}
