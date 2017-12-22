package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// バージョン情報を提供するためのハンドラー
func versionInfo(buildTimeStamp, commit, release string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		info := struct {
			BuildTimestamp string `json:"buildTimestamp"`
			Commit         string `json:"commit"`
			Release        string `json:"release"`
		}{
			buildTimeStamp,
			commit,
			release,
		}

		body, err := json.Marshal(info)
		if err != nil {
			log.Printf("json encoding error: %v", err)
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}
