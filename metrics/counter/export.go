package counter

import (
	"github.com/go-kit/kit/metrics"
	"github.com/weichang-bianjie/metric-sdk/types"
)

type Client interface {
	types.Counter
	GetCounter() metrics.Counter
	MetricName() string
}
