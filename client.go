package metric_sdk

import (
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/weichang-bianjie/metric-sdk/core"
	"github.com/weichang-bianjie/metric-sdk/types"
)

type MetricClient struct {
	pusher *push.Pusher
}

func NewClient(config *types.PusherConfig) MetricClient {
	core.InitPusher(config)
	return MetricClient{}
}

func (client *MetricClient) RegisterMetrics() {
	core.MakeRegister()
	client.pusher = core.GetPusher()
}

func (client *MetricClient) Start() error {
	if err := client.pusher.Add(); err != nil {
		return err
	}
	return nil
}
