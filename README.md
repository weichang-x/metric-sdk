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
 Gauge := gauge.NewGauge
    "",
    "",
    "db_backup_records_processed",
    "The number of records processed in the last DB backup",
    map[string]string{
    "host": "localhost",
    "address":  "127.0.0.1",
    },
)
```

- Counter
```go
 Counter := counter.NewCounter(
    "",
    "",
    "db_backup_records_times_total",
    "The number of records times counter",
    map[string]string{
    "host": "localhost",
    "address":  "127.0.0.1",
    },
)
```

### use in programme
```go

    MetricClient := metric_sdk.NewClient(types.Config{
    Address: ":8080",
    })

	Gauge := gauge.NewGauge
        "",
        "",
        "db_backup_records_processed",
        "The number of records processed in the last DB backup",
        map[string]string{
        "host": "localhost",
        "address":  "127.0.0.1",
        },
	)
	Counter := counter.NewCounter(
        "",
        "",
        "db_backup_records_times_total",
        "The number of records times counter",
        map[string]string{
        "host": "localhost",
        "address":  "127.0.0.1",
        },
	)
	MetricClient.RegisterMetric(Counter, Gauge)
	
	

```
### update the metric value(example: guage)
```go
    guageData := s.Gauge.(types.Guage)
    guageData.Set(float64(1))
    report := func() {
        for {
            t := time.NewTimer(time.Duration(5) * time.Second)
            select {
            case <-t.C:
			 guageData.Add(float64(val))
            
            }
        }
    }
    MetricClient.Start(report)
```