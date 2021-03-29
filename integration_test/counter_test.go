package integration

import (
	"github.com/weichang-bianjie/metric-sdk/types"
	"time"
)

func (s IntegrationTestSuite) TestCounter() {
	cases := []SubTest{
		{
			"TestCounter",
			CounterCase,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func CounterCase(s IntegrationTestSuite) {
	counterData := s.Counter.(types.Counter)
	counterData.Add(2)
	report := func() {
		for {
			t := time.NewTimer(time.Duration(5) * time.Second)
			select {
			case <-t.C:
				counterData.Add(float64(111))
			}
		}
	}
	s.Start(report)

}
