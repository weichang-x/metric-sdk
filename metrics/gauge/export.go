package gauge

import (
	"github.com/weichang-bianjie/metric-sdk/types"
)

type Client interface {
	types.Guage
	Set(value float64)
	MetricName() string
}
