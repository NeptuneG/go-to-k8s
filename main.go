package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/NeptuneG/go-to-k8s/handlers"
	"github.com/NeptuneG/go-to-k8s/version"
)

func main() {
	log.Printf("Starting service ...\ncommit: %s, build timestamp: %s, release: %s",
		version.Commit, version.BuildTimestamp, version.Release)
	port := os.Getenv("APP_PORT")
	if port == "" {
		log.Fatal("port has not been specified")
	}
	router := handlers.Router(version.BuildTimestamp, version.Commit, version.Release)
	interrput := make(chan os.Signal, 1)
	signal.Notify(interrput, os.Interrupt, os.Kill, syscall.SIGTERM)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	go func() {
		log.Fatal(server.ListenAndServe())
	}()
	log.Print("The service is about to serve on :" + port)
	killSignal := <-interrput
	switch killSignal {
	case os.Kill:
		log.Print("Got SIGKILL")
	case os.Interrupt:
		log.Print("Got SIGINT")
	case syscall.SIGTERM:
		log.Print("Got SIGTERM")
	}
	log.Print("The service is about to shut down")
	server.Shutdown(context.Background())
	log.Print("Service shutting down completed")
}
