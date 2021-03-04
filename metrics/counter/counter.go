package counter

import (
	"github.com/golang/protobuf/proto"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
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

func (client clientCounter) NewMetricFamilyScrap(name, help string, groupingLabels map[string]string, value float64) *types.GobbableMetricFamily {
	mf := &types.GobbableMetricFamily{
		Name: proto.String(name),
		Help: proto.String(help),
		Type: dto.MetricType_GAUGE.Enum(),
		Metric: []*dto.Metric{
			{
				Counter: &dto.Counter{
					Value: proto.Float64(value),
				},
			},
		},
	}
	types.SanitizeLabels((*dto.MetricFamily)(mf), groupingLabels)
	return mf
}
