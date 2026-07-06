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
	GetHabits(ctx context.Context, userId uuid.UUID) ([]model.Habit, error)
	CreateHabit(ctx context.Context, userId uuid.UUID, name, description string, isGood bool) (model.Habit, error)
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

func (h *handler) GetHabits(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("user_id")
	userId, err := uuid.Parse(id)

	if err != nil {
		h.logger.Error("invalid user_id", zap.Error(err))
		http.Error(w, "invalid user_id", http.StatusBadRequest)
		return
	}

	habits, err := h.service.GetHabits(r.Context(), userId)

	if err != nil {
		h.logger.Error("failed get habits", zap.Error(err))
		http.Error(w, "failed get habits", http.StatusInternalServerError)
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err = json.NewEncoder(w).Encode(dto.ToHabitsResponse(habits)); err != nil {
		h.logger.Error("failed create response with habits", zap.String("user_id", userId.String()))

	}
}

func (h *handler) CreateHabit(w http.ResponseWriter, r *http.Request) {
	var input dto.HabitRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.logger.Error("failed decode request", zap.Error(err))
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	habit, err := h.service.CreateHabit(r.Context(), input.UserId, input.Name, input.Description, input.IsGood)
	if err != nil {
		h.logger.Error("failed create habit", zap.Error(err))
		http.Error(w, "failed create habit", http.StatusInternalServerError)
	}

	if err = json.NewEncoder(w).Encode(dto.ToHabitResponse(habit)); err != nil {
		h.logger.Error(
			"failed create response with habit",
			zap.String("habit_id", habit.HabitId.String()),
			zap.Error(err),
		)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
