package dto

import (
	"github.com/google/uuid"

	model "github.com/sklyar-vlad/selfDev/internal/model/habit"
)

type HabitResponse struct {
	HabitId     uuid.UUID `json:"habit_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsGood      bool      `json:"is_good"`
}
type HabitsResponse struct {
	Habits []HabitResponse
}

func ToHabitResponse(h model.Habit) HabitResponse {
	return HabitResponse{
		HabitId:     h.HabitId,
		Name:        h.Name,
		Description: h.Description,
		IsGood:      h.IsGood,
	}
}

func ToHabitsResponse(habits []model.Habit) HabitsResponse {
	resp := make([]HabitResponse, 0, len(habits))

	for _, h := range habits {
		resp = append(resp, ToHabitResponse(h))
	}

	return HabitsResponse{
		Habits: resp,
	}
}