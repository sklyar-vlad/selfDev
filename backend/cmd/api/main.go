package main

import (
	"log"
	"net/http"
	"time"

	"github.com/sklyar-vlad/tracker/internal/handler"
)


func main() {
	


	mux := http.NewServeMux()
	handler.RegisterRoutes(mux, userH)

	service := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("service started at localhost:8080")
	log.Fatal(service.ListenAndServe())
}
