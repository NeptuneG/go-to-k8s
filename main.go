package main

import (
	"github.com/NeptuneG/go-to-k8s/handlers"
	"log"
	"net/http"
)

func main() {
	log.Print("Starting service ...")
	router := handlers.Router()
	log.Print("The service is about to serve on :9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}
