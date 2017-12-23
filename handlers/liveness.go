package handlers

import "net/http"

// livenessはk8sヘルスチェックのLiveness Probe
func liveness(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
