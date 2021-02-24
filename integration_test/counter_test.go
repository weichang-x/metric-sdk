package integration

func (s IntegrationTestSuite) TestCounter() {
	cases := []SubTest{
		{
			"TestCounter",
			counterTest,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func counterTest(s IntegrationTestSuite) {
	s.Counter.Add(float64(1))
	s.MetricClient.Start()
}
