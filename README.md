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

    MetricClient := metric_sdk.NewClient(types.Config{
    Address: ":9090",
    })

```

### define metrics labels
- Guage
```go
 Gauge := gauge.NewGauge(
"db_backup",
"records",
"processed",
"The number of records processed in the last DB backup",
[]string{"name","sex"})
```

- Counter
```go
 Counter := counter.NewCounter(
"db_backup",
"records",
"times_total",
"The number of records times counter",
[]string{"name","sex"},
)
```

### use in programme
```go

    MetricClient := metric_sdk.NewClient(types.Config{
    Address: ":8080",
    })

	Gauge := gauge.NewGauge(
	"db_backup",
	"records",
	"processed",
	"The number of records processed in the last DB backup",
	[]string{"name","sex"})
	Counter := counter.NewCounter(
	"db_backup",
	"records",
	"times_total",
	"The number of records times counter",
	[]string{"name","sex"},
	)
	MetricClient.RegisterMetric(Counter, Gauge)
	
	

```
### update the metric value(example: guage)
```go
    guageData := s.Gauge.(gauge.Client)
    report := func() {
        for {
            t := time.NewTimer(time.Duration(5) * time.Second)
            select {
            case <-t.C:
                guageData.With("name","hwc","sex","male").Set(float64(1))
                guageData.With("name","xwd","sex","female").Set(float64(1))
            }
        }
    }
    MetricClient.Start(report)
```