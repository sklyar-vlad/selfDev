package handler

import (
	"net/http"
)

type AuthHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
}

func RegisterRoutes(mux *http.ServeMux, authHandler AuthHandler) {
	mux.HandleFunc("POST /register", authHandler.CreateUser)
	mux.HandleFunc("GET /login", authHandler.GetUser)
}
