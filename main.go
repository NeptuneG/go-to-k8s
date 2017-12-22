package main

import (
	"github.com/NeptuneG/go-to-k8s/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("Starting service ...")
	port := os.Getenv("APP_PORT")
	if port == "" {
		log.Fatal("port has not been specified")
	}
	router := handlers.Router()
	log.Print("The service is about to serve on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
