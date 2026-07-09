package handler

import (
	"net/http"
)

type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	ConfirmEmail(w http.ResponseWriter, r *http.Request)
	Me(w http.ResponseWriter, r *http.Request)
}

// TODO: GetUsers(w http.ResponseWriter, r *http.Request)
// TODO: DeleteUser(w http.ResponseWriter, r *http.Request)
// TODO: UpdateUser(w http.ResponseWriter, r *http.Request)
type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
}

// TODO: UpdateHabit(w http.ResponseWriter, r *http.Request)
type HabitHandler interface {
	GetHabits(w http.ResponseWriter, r *http.Request)
	CreateHabit(w http.ResponseWriter, r *http.Request)
	DeleteHabit(w http.ResponseWriter, r *http.Request)
	ConfirmHabit(w http.ResponseWriter, r *http.Request)
	CancelHabit(w http.ResponseWriter, r *http.Request)
	GetHabitConfirmDates(w http.ResponseWriter, r *http.Request)
}

// TODO: mux.HandleFunc("GET /api/users", userHandler.GetUsers)
// TODO: mux.HandleFunc("PATCH /api/users/{id}", userHandler.UpdateUser)
// TODO: mux.HandleFunc("DELETE /api/users/{id}", userHandler.DeleteUser)
// TODO: mux.HandleFunc("PATCH /api/habit/{id}", habitHandler.UpdateHabit)
func RegisterRoutes(mux *http.ServeMux, userHandler UserHandler, authHandler AuthHandler, habitHandler HabitHandler) {
	mux.HandleFunc("POST /api/users", userHandler.CreateUser)
	mux.HandleFunc("GET /api/users/{id}", userHandler.GetUser)

	mux.HandleFunc("GET /api/auth/me", authHandler.Me)
	mux.HandleFunc("POST /api/login", authHandler.Login)
	mux.HandleFunc("POST /api/logout", authHandler.Logout)
	mux.HandleFunc("POST /api/refresh", authHandler.Refresh)
	mux.HandleFunc("POST /api/register", authHandler.Register)
	mux.HandleFunc("POST /api/verify/{token}", authHandler.ConfirmEmail)

	mux.HandleFunc("GET /api/habit/{user_id}", habitHandler.GetHabits)
	mux.HandleFunc("POST /api/habit", habitHandler.CreateHabit)
	mux.HandleFunc("DELETE /api/habit/{id}", habitHandler.DeleteHabit)
	mux.HandleFunc("POST /api/habit/{id}/confirm", habitHandler.ConfirmHabit)
	mux.HandleFunc("POST /api/habit/{id}/cancel", habitHandler.CancelHabit)
	mux.HandleFunc("GET /api/habit/{id}/confirmed", habitHandler.GetHabitConfirmDates)
}
