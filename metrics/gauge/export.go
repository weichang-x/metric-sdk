package gauge

import (
	"github.com/go-kit/kit/metrics"
	"github.com/weichang-bianjie/metric-sdk/types"
)

type Client interface {
	types.Guage
	GetGuage() metrics.Gauge
	MetricName() string
}
