package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/sklyar-vlad/tracker/backend/docs"
	"github.com/sklyar-vlad/tracker/backend/internal/handler"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	mux := http.NewServeMux()
	habitHandler := &handler.HabitHandler{}
	mux.Handle("/habit", habitHandler)
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	service := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("service started at localhost:8080")
	log.Fatal(service.ListenAndServe())
}
