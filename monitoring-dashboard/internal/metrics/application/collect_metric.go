package application

import (
	"monitoring-dashboard/internal/metrics/domain"
	"time"
)

type MetricService struct {
	repo domain.MetricRepository
}

func NewMetricService(repo domain.MetricRepository) *MetricService {
	return &MetricService{repo: repo}
}

func (s *MetricService) CollectMetric(name string, value float64) error {
	metric := domain.Metric{
		Name:      name,
		Value:     value,
		Timestamp: time.Now(),
	}
	return s.repo.Save(metric)
}
