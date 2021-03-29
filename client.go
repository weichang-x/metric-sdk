package metric_sdk

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/weichang-bianjie/metric-sdk/metrics"
	"log"
	"net/http"
)

type MetricClient interface {
	RegisterMetric(metrics ...metrics.Metric)
	Start(report func())
}

type metricClient struct {
	metricsProvider map[string]metrics.Metric
}

func NewClient() MetricClient {
	return metricClient{
		metricsProvider: make(map[string]metrics.Metric, 1),
	}
}

func (m metricClient) RegisterMetric(metrics ...metrics.Metric) {
	for _, one := range metrics {
		m.metricsProvider[one.MetricName()] = one
	}
}
func (m metricClient) Start(report func()) {
	go report()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: promhttp.Handler(),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	select {}
}
