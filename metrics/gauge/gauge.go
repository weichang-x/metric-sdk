package gauge

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/weichang-bianjie/metric-sdk/core"
	"github.com/weichang-bianjie/metric-sdk/types"
)

func NewGuage() Client {
	return clientGuage{}
}

type clientGuage struct {
}

func (client clientGuage) RegisterMetricInfo(name string, help string, constLabels map[string]string) types.Gauge {
	completionOpts := prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        name,
		Help:        help,
		ConstLabels: constLabels,
	})
	core.RegisterCollector(completionOpts)
	return completionOpts
}
