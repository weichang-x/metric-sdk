package metric_sdk

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/weichang-bianjie/metric-sdk/metrics"
	"github.com/weichang-bianjie/metric-sdk/types"
	"log"
	"net/http"
)

type MetricClient interface {
	RegisterMetric(metrics ...metrics.Metric)
	Start(report func())
}

type metricClient struct {
	cfg             types.Config
	metricsProvider map[string]metrics.Metric
}

func NewClient(config types.Config) MetricClient {
	return metricClient{
		metricsProvider: make(map[string]metrics.Metric, 1),
		cfg:             config,
	}
}

// Deprecated: RegisterMetric defines for check metrics if exist by metrics name,but no use
func (m metricClient) RegisterMetric(metrics ...metrics.Metric) {
	for _, one := range metrics {
		m.metricsProvider[one.MetricName()] = one
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
