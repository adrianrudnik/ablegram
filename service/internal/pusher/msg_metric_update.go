package pusher

type MetricUpdatePush struct {
	Type   string            `json:"type"`
	Values map[string]uint64 `json:"values"`
}

func NewMetricUpdatePush(values map[string]uint64) *MetricUpdatePush {
	return &MetricUpdatePush{
		Type:   "metric_update",
		Values: values,
	}
}
