package core

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/promlog"
	"github.com/prometheus/common/route"
	"github.com/prometheus/common/version"
	"github.com/prometheus/exporter-toolkit/web"
	"github.com/weichang-bianjie/metric-sdk/types"
	"log"
	"net"
	"net/http"
	"time"
)

var (
	httpCnt = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "scrap_http_requests_total",
			Help: "Total HTTP requests processed by the Scrap",
		},
		[]string{"handler", "code", "method"},
	)
	familyMaps []types.MetricGroup
)

func NewHttpServer(address string) error {
	r := route.New()
	r.Get("/metrics", scrapFunc("api/v1/metrics", metrics))
	r.Get("/status", scrapFunc("api/v1/status", status))
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	logger := promlog.New(&promlog.Config{})
	err = web.Serve(l, &http.Server{Addr: address, Handler: r}, "", logger)

	return err
}

func scrapFunc(handlerName string, f http.HandlerFunc) http.HandlerFunc {
	return instrumentWithCounter(
		handlerName,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			f(w, r)
		}),
	)
}

func instrumentWithCounter(handlerName string, handler http.Handler) http.HandlerFunc {
	return promhttp.InstrumentHandlerCounter(
		httpCnt.MustCurryWith(prometheus.Labels{"handler": handlerName}),
		handler,
	)
}

func RegisterScrapMetric(group ...types.MetricGroup) {
	familyMaps = append(familyMaps, group...)
}

func status(w http.ResponseWriter, r *http.Request) {
	res := map[string]interface{}{}
	flags := map[string]string{}
	res["flags"] = flags
	res["start_time"] = time.Now()
	res["build_information"] = map[string]string{
		"version":   version.Version,
		"revision":  version.Revision,
		"branch":    version.Branch,
		"buildUser": version.BuildUser,
		"buildDate": version.BuildDate,
		"goVersion": version.GoVersion,
	}

	respond(w, res)
}

func metrics(w http.ResponseWriter, r *http.Request) {
	res := []interface{}{}
	for _, v := range familyMaps {
		metricResponse := map[string]interface{}{}
		metricResponse["labels"] = v.Labels
		for _, value := range v.Metrics {
			metricFamily := (*dto.MetricFamily)(value)
			metricResponse[*value.Name] = types.Metrics{
				Type:      metricFamily.GetType().String(),
				Help:      metricFamily.GetHelp(),
				Timestamp: time.Now(),
				Metrics:   makeScrapMetrics(metricFamily.GetMetric(), metricFamily.GetType()),
			}
		}
		res = append(res, metricResponse)
	}
	respond(w, res)
}

func makeScrapMetrics(metrics []*dto.Metric, metricsType dto.MetricType) []types.ScrapMetric {

	jsonMetrics := make([]types.ScrapMetric, len(metrics))

	for i, m := range metrics {

		metric := types.ScrapMetric{}
		metric["labels"] = makeLabels(m)
		switch metricsType {
		case dto.MetricType_SUMMARY:
			metric["quantiles"] = makeQuantiles(m)
			metric["count"] = fmt.Sprint(m.GetSummary().GetSampleCount())
			metric["sum"] = fmt.Sprint(m.GetSummary().GetSampleSum())
		case dto.MetricType_HISTOGRAM:
			metric["buckets"] = makeBuckets(m)
			metric["count"] = fmt.Sprint(m.GetHistogram().GetSampleCount())
			metric["sum"] = fmt.Sprint(m.GetHistogram().GetSampleSum())
		default:
			metric["value"] = fmt.Sprint(getValue(m))
		}
		jsonMetrics[i] = metric
	}
	return jsonMetrics
}

func makeLabels(m *dto.Metric) map[string]string {
	result := map[string]string{}
	for _, lp := range m.Label {
		result[lp.GetName()] = lp.GetValue()
	}
	return result
}

func makeQuantiles(m *dto.Metric) map[string]string {
	result := map[string]string{}
	for _, q := range m.GetSummary().Quantile {
		result[fmt.Sprint(q.GetQuantile())] = fmt.Sprint(q.GetValue())
	}
	return result
}

func makeBuckets(m *dto.Metric) map[string]string {
	result := map[string]string{}
	for _, b := range m.GetHistogram().Bucket {
		result[fmt.Sprint(b.GetUpperBound())] = fmt.Sprint(b.GetCumulativeCount())
	}
	return result
}

func getValue(m *dto.Metric) float64 {
	switch {
	case m.Gauge != nil:
		return m.GetGauge().GetValue()
	case m.Counter != nil:
		return m.GetCounter().GetValue()
	case m.Untyped != nil:
		return m.GetUntyped().GetValue()
	default:
		return 0
	}
}

type respstatus string

const (
	statusSuccess respstatus = "success"
	statusError   respstatus = "error"
)

type errorType string

const (
	errorNone        errorType = ""
	errorTimeout     errorType = "timeout"
	errorCanceled    errorType = "canceled"
	errorExec        errorType = "execution"
	errorBadData     errorType = "bad_data"
	errorInternal    errorType = "internal"
	errorUnavailable errorType = "unavailable"
	errorNotFound    errorType = "not_found"
)

type response struct {
	Status    respstatus  `json:"status"`
	Data      interface{} `json:"data,omitempty"`
	ErrorType errorType   `json:"errorType,omitempty"`
	Error     string      `json:"error,omitempty"`
}

func respond(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	b, err := json.Marshal(&response{
		Status: statusSuccess,
		Data:   data,
	})
	if err != nil {
		log.Println("error marshaling JSON", "err", err)
		respondError(w, apiError{
			typ: errorBadData,
			err: err,
		}, "")
	}

	if _, err := w.Write(b); err != nil {
		log.Println("failed to write data to connection", "err", err)
	}
}

type apiError struct {
	typ errorType
	err error
}

func (e *apiError) Error() string {
	return fmt.Sprintf("%s: %s", e.typ, e.err)
}

func respondError(w http.ResponseWriter, apiErr apiError, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	switch apiErr.typ {
	case errorBadData:
		w.WriteHeader(http.StatusBadRequest)
	case errorInternal:
		w.WriteHeader(http.StatusInternalServerError)
	default:
		panic(fmt.Sprintf("unknown error type %q", apiErr.Error()))
	}

	b, err := json.Marshal(&response{
		Status:    statusError,
		ErrorType: apiErr.typ,
		Error:     apiErr.err.Error(),
		Data:      data,
	})
	if err != nil {
		return
	}
	log.Println("API error", "err", apiErr.Error())

	if _, err := w.Write(b); err != nil {
		log.Println("failed to write data to connection", "err", err)
	}
}
