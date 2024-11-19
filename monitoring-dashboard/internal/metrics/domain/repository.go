package domain

type MetricRepository interface {
	Save(metric Metric) error
	FindAll() ([]Metric, error)
}
