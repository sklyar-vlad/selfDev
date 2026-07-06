package habit

import "github.com/google/uuid"

type Habit struct {
	HabitId     uuid.UUID
	UserId      uuid.UUID
	Name        string
	Description string
	IsGood      bool
}

func NewHabit(userId uuid.UUID, name, description string, isGood bool) (Habit, error) {
	return Habit{uuid.New(), userId, name, description, isGood}, nil
}
