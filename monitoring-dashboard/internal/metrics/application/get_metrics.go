package application

import "monitoring-dashboard/internal/metrics/domain"

func (s *MetricService) GetAllMetrics() ([]domain.Metric, error) {
	return s.repo.FindAll()
}
