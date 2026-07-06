package dto

import "github.com/google/uuid"

type HabitRequest struct {
	UserId      uuid.UUID `json:"username"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
}
