package persistence

import (
	"database/sql"
	"log"
	"monitoring-dashboard/internal/metrics/domain"
)

type PostgresMetricRepository struct {
	db     *sql.DB
	logger *log.Logger
}

func NewPostgresMetricRepository(db *sql.DB, logger *log.Logger) *PostgresMetricRepository {
	return &PostgresMetricRepository{db: db, logger: logger}
}

func (r *PostgresMetricRepository) Save(metric domain.Metric) error {
	_, err := r.db.Exec("INSERT INTO metrics (name, value, timestamp) VALUES ($1, $2, $3)", metric.Name, metric.Value, metric.Timestamp)
	return err
}

func (r *PostgresMetricRepository) FindAll() ([]domain.Metric, error) {
	rows, err := r.db.Query("SELECT name, value, timestamp FROM metrics")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var metrics []domain.Metric
	for rows.Next() {
		var metric domain.Metric
		if err := rows.Scan(&metric.Name, &metric.Value, &metric.Timestamp); err != nil {
			return nil, err
		}
		metrics = append(metrics, metric)
	}

	return metrics, nil
}
