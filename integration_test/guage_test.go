package integration

import (
	"github.com/weichang-bianjie/metric-sdk/types"
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
	guageData := s.Gauge.(types.Guage)
	report := func() {
		for {
			t := time.NewTimer(time.Duration(5) * time.Second)
			select {
			case <-t.C:
				guageData.Set(float64(1111))
			}
		}
	}
	s.Start(report)

}
