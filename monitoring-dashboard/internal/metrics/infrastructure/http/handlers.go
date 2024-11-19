package http

import (
	"encoding/json"
	"log"
	"monitoring-dashboard/internal/metrics/application"
	"net/http"

	"github.com/gorilla/mux"
)

type MetricHandler struct {
	service *application.MetricService
}

func NewMetricHandler(service *application.MetricService) *MetricHandler {
	return &MetricHandler{service: service}
}

func (h *MetricHandler) Collect(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name  string  `json:"name"`
		Value float64 `json:"value"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.service.CollectMetric(req.Name, req.Value); err != nil {
		http.Error(w, "Failed to save metric", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *MetricHandler) GetMetrics(w http.ResponseWriter, r *http.Request) {
	metrics, err := h.service.GetAllMetrics()
	if err != nil {
		http.Error(w, "Failed to retrieve metrics", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(metrics)
}

func RegisterHandlers(router *mux.Router, service *application.MetricService, logger *log.Logger) {
	handler := NewMetricHandler(service)

	router.HandleFunc("/metrics/db", handler.Collect).Methods("POST")
	router.HandleFunc("/metrics/db", handler.GetMetrics).Methods("GET")
}
