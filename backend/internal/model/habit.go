package model

import (
	"time"

	"uuid"
)

type Habit struct {
	habit_id    uuid.UUID
	user_id     uuid.UUID
	name        string
	description string
	habit_type  string
	created_at  time.Time
}

type Habits_completed struct {
	id           uuid.UUID
	habit_id     uuid.UUID
	completed_at time.Time
}
