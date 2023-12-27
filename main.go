package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"go-initializr/common"
	"go-initializr/service"
)

type Config struct {
	service *service.Service
}

func main() {

	app := Config{}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.routes(),
	}

	common.Write()
	common.Remove()

	log.Printf("Starting modules service on port %s...", port)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
