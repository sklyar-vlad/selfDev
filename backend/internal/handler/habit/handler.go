package habit

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/sklyar-vlad/selfDev/internal/handler/habit/dto"
	model "github.com/sklyar-vlad/selfDev/internal/model/habit"
)

type HabitService interface {
	// GetHabits(ctx context.Context, userId uuid.UUID) ([]model.Habit, error)
	CreateHabit(ctx context.Context, userId uuid.UUID, name, description, category string) (model.Habit, error)
}

type handler struct {
	service HabitService
	logger  *zap.Logger
}

func NewHandler(service HabitService, logger *zap.Logger) *handler {
	return &handler{
		service: service,
		logger:  logger,
	}
}

// func (h *handler) GetHabits(w http.ResponseWriter, r *http.Request) {
// 	return
// }

// func (h *handler) Register(w http.ResponseWriter, r *http.Request) {
// 	var input dto.AuthRequest

// 	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
// 		h.logger.Error("failed decode request", zap.Error(err))
// 		http.Error(w, "invalid json", http.StatusBadRequest)
// 		return
// 	}

// 	if err := h.service.Register(r.Context(), input.Username, input.Email, input.Password); err != nil {
// 		h.logger.Error("failed create user", zap.Error(err))
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// }

func (h *handler) CreateHabit(w http.ResponseWriter, r *http.Request) {
	var input dto.HabitRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.logger.Error("failed decode request", zap.Error(err))
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	habit, err := h.service.CreateHabit(r.Context(), input.UserId, input.Name, input.Description, input.Category)

	if err != nil {
		h.logger.Error("failed create habit", zap.Error(err))
		http.Error(w, "failed create habit", http.StatusInternalServerError)
	}

	if err = json.NewEncoder(w).Encode(dto.ToHabitResponse(habit)); err != nil {
		h.logger.Error("failed create response with habit", zap.String("habit_id", habit.HabitId.String()), zap.Error(err))
	}
}
