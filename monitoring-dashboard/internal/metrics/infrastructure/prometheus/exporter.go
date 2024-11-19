package prometheus

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// ExportMetricsHandler exposes the metrics via HTTP
func ExportMetricsHandler() http.Handler {
	return promhttp.Handler()
}
