package domain

import "time"

type Metric struct {
	Name      string
	Value     float64
	Timestamp time.Time
}

func NewMetric(name string, value float64) (*Metric, error) {
	if name == "" {
		return nil, ErrMetricNotFound
	}
	return &Metric{Name: name, Value: value, Timestamp: time.Now()}, nil
}
