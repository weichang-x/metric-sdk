package counter

import "github.com/weichang-bianjie/metric-sdk/types"

type Client interface {
	RegisterMetricInfo(name string, help string, constLabels map[string]string) types.Counter
}
