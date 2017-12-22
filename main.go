package main

import (
	"github.com/NeptuneG/go-to-k8s/handlers"
	"github.com/NeptuneG/go-to-k8s/version"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Printf("Starting service ...\ncommit: %s, build timestamp: %s, release: %s",
		version.Commit, version.BuildTimestamp, version.Release)
	port := os.Getenv("APP_PORT")
	if port == "" {
		log.Fatal("port has not been specified")
	}
	router := handlers.Router(version.BuildTimestamp, version.Commit, version.Release)
	log.Print("The service is about to serve on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
