package integration

import (
	"github.com/stretchr/testify/suite"
	metric_sdk "github.com/weichang-bianjie/metric-sdk"
	"github.com/weichang-bianjie/metric-sdk/core"
	"github.com/weichang-bianjie/metric-sdk/metrics/counter"
	"github.com/weichang-bianjie/metric-sdk/metrics/gauge"
	"github.com/weichang-bianjie/metric-sdk/types"
	"testing"
)

type IntegrationTestSuite struct {
	metric_sdk.MetricClient
	Counter types.Counter
	Gauge   types.Gauge
	suite.Suite
}

type SubTest struct {
	testName string
	testCase func(s IntegrationTestSuite)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.MetricClient = metric_sdk.NewClient(&types.PusherConfig{
		PushGatewayUrl: "http://localhost:9091",
		JobName:        "prometheus",
	})
	core.RegisterGroupLabel("name", "hwc")
	core.RegisterGroupLabel("host", "127.0.0.1")
	s.Gauge = gauge.NewGuage().RegisterMetricInfo(
		"db_backup_records_processed",
		"The number of records processed in the last DB backup",
		map[string]string{
			"ip": "127.0.0.1",
		},
	)
	s.Counter = counter.NewCounter().RegisterMetricInfo(
		"db_backup_records_times_total",
		"The number of records times counter",
		map[string]string{
			"ip": "127.0.0.1",
		},
	)
	s.MetricClient.RegisterMetrics()
}
