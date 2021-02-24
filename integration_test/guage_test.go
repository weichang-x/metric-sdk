package integration

import "time"

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
	s.Gauge.Set(float64(time.Now().Unix()))
	s.MetricClient.Start()
}
