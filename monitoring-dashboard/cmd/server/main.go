package main

import (
	"database/sql"
	"monitoring-dashboard/internal/metrics/application"
	"monitoring-dashboard/internal/metrics/infrastructure/http"
	"monitoring-dashboard/internal/metrics/infrastructure/persistence"
	"monitoring-dashboard/pkg/config"
	"monitoring-dashboard/pkg/logging"
	https "net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Setup logger
	logger := logging.NewLogger()

	// Connect to the database
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize the metric repository
	repo := persistence.NewPostgresMetricRepository(db, logger)

	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}

	service := application.NewMetricService(repo)

	// Register metrics
	metricCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "metric_collection_total",
			Help: "Total number of collected metrics",
		},
		[]string{"name"},
	)
	prometheus.MustRegister(metricCounter)

	// HTTP router
	router := mux.NewRouter()

	// Metrics endpoint for Prometheus scraping
	router.Handle("/metrics", promhttp.Handler()) // Exposes the metrics at /metrics

	// API endpoints
	http.RegisterHandlers(router, service, logger)

	// Start HTTP server
	logger.Println("Starting server on", cfg.ServerPort)
	if err := https.ListenAndServe(cfg.ServerPort, router); err != nil {
		logger.Fatalf("Server failed: %v", err)
	}
}
