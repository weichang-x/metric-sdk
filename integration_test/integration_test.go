package integration

import (
	"github.com/stretchr/testify/suite"
	metric_sdk "github.com/weichang-bianjie/metric-sdk"
	"github.com/weichang-bianjie/metric-sdk/metrics"
	"github.com/weichang-bianjie/metric-sdk/metrics/counter"
	"github.com/weichang-bianjie/metric-sdk/metrics/gauge"
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
	s.MetricClient = metric_sdk.NewClient()

	s.Gauge = gauge.NewGauge(
		"",
		"",
		"db_backup_records_processed",
		"The number of records processed in the last DB backup",
		[]string{},
	)
	s.Counter = counter.NewCounter(
		"",
		"",
		"db_backup_records_times_total",
		"The number of records times counter",
		[]string{},
	)
	s.RegisterMetric(s.Counter, s.Gauge)

}
