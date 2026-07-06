package dto

import (
	"github.com/google/uuid"

	model "github.com/sklyar-vlad/selfDev/internal/model/habit"
)

type HabitResponse struct {
	HabitId uuid.UUID `json:"habit_id"`
}

func ToHabitResponse(p model.Habit) HabitResponse {
	resp := HabitResponse{
		HabitId: p.HabitId,
	}

	return resp
}
