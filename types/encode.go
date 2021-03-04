package types

import (
	"github.com/golang/protobuf/proto"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/model"
	"sort"
)

// GobbableMetricFamily is a dto.MetricFamily that implements GobDecoder and
// GobEncoder.
type GobbableMetricFamily dto.MetricFamily

// GobDecode implements gob.GobDecoder.
func (gmf *GobbableMetricFamily) GobDecode(b []byte) error {
	return proto.Unmarshal(b, (*dto.MetricFamily)(gmf))
}

// GobEncode implements gob.GobEncoder.
func (gmf *GobbableMetricFamily) GobEncode() ([]byte, error) {
	return proto.Marshal((*dto.MetricFamily)(gmf))
}

// sanitizeLabels ensures that all the labels in groupingLabels and the
// `instance` label are present in the MetricFamily. The label values from
// groupingLabels are set in each Metric, no matter what. After that, if the
// 'instance' label is not present at all in a Metric, it will be created (with
// an empty string as value).
//
// Finally, sanitizeLabels sorts the label pairs of all metrics.
func SanitizeLabels(mf *dto.MetricFamily, groupingLabels map[string]string) {
	gLabelsNotYetDone := make(map[string]string, len(groupingLabels))

metric:
	for _, m := range mf.GetMetric() {
		for ln, lv := range groupingLabels {
			gLabelsNotYetDone[ln] = lv
		}
		hasInstanceLabel := false
		for _, lp := range m.GetLabel() {
			ln := lp.GetName()
			if lv, ok := gLabelsNotYetDone[ln]; ok {
				lp.Value = proto.String(lv)
				delete(gLabelsNotYetDone, ln)
			}
			if ln == string(model.InstanceLabel) {
				hasInstanceLabel = true
			}
			if len(gLabelsNotYetDone) == 0 && hasInstanceLabel {
				sort.Sort(labelPairs(m.Label))
				continue metric
			}
		}
		for ln, lv := range gLabelsNotYetDone {
			m.Label = append(m.Label, &dto.LabelPair{
				Name:  proto.String(ln),
				Value: proto.String(lv),
			})
			if ln == string(model.InstanceLabel) {
				hasInstanceLabel = true
			}
			delete(gLabelsNotYetDone, ln) // To prepare map for next metric.
		}
		if !hasInstanceLabel {
			m.Label = append(m.Label, &dto.LabelPair{
				Name:  proto.String(string(model.InstanceLabel)),
				Value: proto.String(""),
			})
		}
		sort.Sort(labelPairs(m.Label))
	}
}

// labelPairs implements sort.Interface. It provides a sortable version of a
// slice of dto.LabelPair pointers.
type labelPairs []*dto.LabelPair

func (s labelPairs) Len() int {
	return len(s)
}

func (s labelPairs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s labelPairs) Less(i, j int) bool {
	return s[i].GetName() < s[j].GetName()
}
