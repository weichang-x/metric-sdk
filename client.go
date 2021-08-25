package metric_sdk

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/weichang-bianjie/metric-sdk/types"
	"log"
	"net/http"
)

type MetricClient interface {
	Start(report func())
}

type metricClient struct {
	cfg types.Config
}

func NewClient(config types.Config) MetricClient {
	return metricClient{
		cfg: config,
	}
}

func (m metricClient) Start(report func()) {
	go report()
	srv := &http.Server{
		Addr:    m.cfg.Address,
		Handler: promhttp.Handler(),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	select {}
}
