package handler

import (
	"net/http"
)

type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	ConfirmEmail(w http.ResponseWriter, r *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
}

type UserHandler interface {
	// GetUsers(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	// DeleteUser(w http.ResponseWriter, r *http.Request)
	// UpdateUser(w http.ResponseWriter, r *http.Request)
}

type HabitHandler interface {
	GetHabits(w http.ResponseWriter, r *http.Request)
	// CreateHabit(w http.ResponseWriter, r *http.Request)
	// GetHabit(w http.ResponseWriter, r *http.Request)
	// DeleteHabit(w http.ResponseWriter, r *http.Request)
	// UpdateHabit(w http.ResponseWriter, r *http.Request)
}

func RegisterRoutes(mux *http.ServeMux, userHandler UserHandler, authHandler AuthHandler, habitHandler HabitHandler) {
	// User
	mux.HandleFunc("POST /api/users", userHandler.CreateUser)
	// mux.HandleFunc("GET /api/users", userHandler.GetUsers)
	mux.HandleFunc("GET /api/users/{id}", userHandler.GetUser)
	// mux.HandleFunc("PATCH /api/users/{id}", userHandler.UpdateUser)
	// mux.HandleFunc("DELETE /api/users/{id}", userHandler.DeleteUser)

	// Auth
	mux.HandleFunc("POST /api/login", authHandler.Login)
	mux.HandleFunc("POST /api/logout", authHandler.Logout)
	mux.HandleFunc("POST /api/register", authHandler.Register)
	mux.HandleFunc("POST /api/verify/{token}", authHandler.ConfirmEmail)
	mux.HandleFunc("POST /api/refresh", authHandler.Refresh)

	// Habit
	// mux.HandleFunc("POST /api/habit", habitHandler.CreateHabit)
	mux.HandleFunc("GET /api/habit", habitHandler.GetHabits)
	// mux.HandleFunc("GET /api/habit/{id}", habitHandler.GetHabit)
	// mux.HandleFunc("PATCH /api/habit/{id}", habitHandler.UpdateHabit)
	// mux.HandleFunc("DELETE /api/habit/{id}", habitHandler.DeleteHabit)
}
