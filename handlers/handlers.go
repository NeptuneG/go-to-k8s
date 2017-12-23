package handlers

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/gorilla/mux"
)

/*
   だくさんのルートルールを対応するため、
   ルータの初期化とルールをpackageに含まれる
*/

// 必要なルートを登録し、ルータのインスタンスを返す
func Router(buildTimestamp, commit, release string) *mux.Router {
	isReady := &atomic.Value{}
	isReady.Store(false)
	go func() {
		log.Printf("readiness probe is set to be negative")
		time.Sleep(3 * time.Second)
		isReady.Store(true)
		log.Printf("readiness probe is positive now.")
	}()
	r := mux.NewRouter()
	r.HandleFunc("/home", home).Methods("GET")
	r.HandleFunc("/version", versionInfo(buildTimestamp, commit, release)).Methods("GET")
	r.HandleFunc("/liveness", liveness)
	r.HandleFunc("/readiness", readiness(isReady))
	return r
}
