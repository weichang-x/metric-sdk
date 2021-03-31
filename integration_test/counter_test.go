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
	report := func() {
		for {
			t := time.NewTimer(time.Duration(5) * time.Second)
			select {
			case <-t.C:
				for _, val := range []int64{1, 2, 3, 4, 5} {
					counterData.Add(float64(val))
				}

			}
		}
	}
	s.Start(report)

}
