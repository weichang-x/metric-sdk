package integration

import (
	"github.com/weichang-bianjie/metric-sdk/metrics/gauge"
	"time"
)

func (s IntegrationTestSuite) TestGauge() {
	cases := []SubTest{
		{
			"TestGauge",
			guage,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func guage(s IntegrationTestSuite) {
	guageData := s.Gauge.(gauge.Client)
	report := func() {
		for {
			t := time.NewTimer(time.Duration(5) * time.Second)
			select {
			case <-t.C:
				guageData.With("name", "hwc", "sex", "male").Set(float64(1))
				guageData.With("name", "xwd", "sex", "female").Set(float64(1))
			}
		}
	}
	s.Start(report)

}
