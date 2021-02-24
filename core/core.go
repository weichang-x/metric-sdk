package core

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/weichang-bianjie/metric-sdk/types"
)

var (
	pusher      *push.Pusher
	collectors  []prometheus.Collector
	groupLabels = make(map[string]string, 1)
)

func InitPusher(config *types.PusherConfig) {
	pusher = push.New(config.PushGatewayUrl, config.JobName)
}

func MakeRegister() {
	registry := prometheus.NewRegistry()
	registry.MustRegister(collectors...)
	pusher = pusher.Gatherer(registry)
	for name, value := range groupLabels {
		pusher = pusher.Grouping(name, value)
	}
}

func GetPusher() *push.Pusher {
	return pusher
}

func RegisterCollector(cs ...prometheus.Collector) {
	collectors = append(collectors, cs...)
}

func RegisterGroupLabel(name, value string) {
	if _, ok := groupLabels[name]; !ok {
		groupLabels[name] = value
	}
}
