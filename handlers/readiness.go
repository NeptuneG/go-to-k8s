package handlers

import (
	"net/http"
	"sync/atomic"
)

func readiness(isReady *atomic.Value) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		if isReady == nil || !isReady.Load().(bool) {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		}
		w.WriteHeader(http.StatusOK)
	}
}
