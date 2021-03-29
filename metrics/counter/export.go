package counter

import "github.com/weichang-bianjie/metric-sdk/types"

type Client interface {
	types.Counter
	Add(value float64)
	MetricName() string
}
