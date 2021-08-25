package integration

import (
	"github.com/stretchr/testify/suite"
	metric_sdk "github.com/weichang-bianjie/metric-sdk"
	"github.com/weichang-bianjie/metric-sdk/metrics"
	"github.com/weichang-bianjie/metric-sdk/metrics/counter"
	"github.com/weichang-bianjie/metric-sdk/metrics/gauge"
	"github.com/weichang-bianjie/metric-sdk/types"
	"testing"
)

type IntegrationTestSuite struct {
	metric_sdk.MetricClient
	Counter metrics.Metric
	Gauge   metrics.Metric
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
	s.MetricClient = metric_sdk.NewClient(types.Config{
		Address: ":8080",
	})

	s.Gauge = gauge.NewGauge(
		"db_backup",
		"records",
		"processed",
		"The number of records processed in the last DB backup",
		[]string{"name", "sex"})
	s.Counter = counter.NewCounter(
		"db_backup",
		"records",
		"times_total",
		"The number of records times counter",
		[]string{"name", "sex"},
	)

}
