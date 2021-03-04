package integration

import (
	"fmt"
	"github.com/prometheus/common/promlog"
	"github.com/prometheus/exporter-toolkit/web"
	"github.com/weichang-bianjie/metric-sdk/core"
	"github.com/weichang-bianjie/metric-sdk/types"
	"net"
	"net/http"
	"os"
)

func (s IntegrationTestSuite) TestCounter() {
	cases := []SubTest{
		{
			"TestCounter",
			counterTest,
		},
		{
			"TestCounter",
			counterScrap,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func counterTest(s IntegrationTestSuite) {
	//s.Counter.Add(float64(1))
	//s.MetricClient.Start()
}

func counterScrap(s IntegrationTestSuite) {
	labels := map[string]string{
		"ip":   "127.0.0.1",
		"name": "hwc",
		"sex":  "male",
	}

	metricGroup := types.MetricGroup{
		Labels: labels,
		Metrics: []*types.GobbableMetricFamily{
			s.CountClient.NewMetricFamilyScrap(
				"scrap_backup_records_times_total",
				"The number of records times scrap counter",
				map[string]string{
					"job":      "sdk-metric",
					"instance": "localhost",
				}, 1,
			),
			s.GaugeClient.NewMetricFamilyScrap(
				"scrap_backup_records_times",
				"The number of records times scrap gauge",
				map[string]string{
					"job":      "sdk-metric",
					"instance": "localhost",
				}, 1,
			),
		},
	}
	core.RegisterScrapMetric(metricGroup)
	r := core.NewRoute()
	l, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	logger := promlog.New(&promlog.Config{})
	err = web.Serve(l, &http.Server{Addr: "localhost:8888", Handler: r}, "", logger)

	if err != nil {
		fmt.Println(err.Error())
	}
}
