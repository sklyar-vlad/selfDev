package habit

import "github.com/google/uuid"

type Habit struct {
	HabitId     uuid.UUID
	UserId      uuid.UUID
	Name        string
	Description string
	Category    string
}

func NewHabit(userId uuid.UUID, name, description, category string) (Habit, error) {
	return Habit{uuid.New(), userId, name, description, category}, nil
}
