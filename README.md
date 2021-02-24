# metric-sdk
client prometheus metric sdk

## install

### Requirement

Go version above 1.15

### Use Go Mod

```text
require (
    github.com/weichang-bianjie/metric-sdk latest
)
```

## Usage

### Init Client

The initialization SDK code is as follows:

```go

	s.MetricClient = metric_sdk.NewClient(&types.PusherConfig{
		PushGatewayUrl: "http://localhost:9091",
		JobName:        "prometheus",
	})
```

define group labels
```go
 core.RegisterGroupLabel("name", "xiaoming")
 core.RegisterGroupLabel("host", "127.0.0.1")
```

use in programme
```go

   MetricClient = metric_sdk.NewClient(&types.PusherConfig{
		PushGatewayUrl: "http://localhost:9091",
		JobName:        "prometheus",
	})
	core.RegisterGroupLabel("name", "xiaoming")
	core.RegisterGroupLabel("host", "127.0.0.1")
	Gauge = gauge.NewGuage().RegisterMetricInfo(
		"db_backup_records_processed",
		"The number of records processed in the last DB backup",
		map[string]string{
			"ip": "127.0.0.1",
		},
	)
	Counter = counter.NewCounter().RegisterMetricInfo(
		"db_backup_records_times_total",
		"The number of records times counter",
		map[string]string{
			"ip": "127.0.0.1",
		},
	)
	MetricClient.RegisterMetrics()

```
update the metric value
```go
    Counter.Add(float64(2))
 	MetricClient.Start()
```