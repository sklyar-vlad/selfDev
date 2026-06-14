package handler

import "github.com/sklyar-vlad/tracker/backend/internal/model"

type HabitService interface {
	GetAllHabit() []model.Habit
}
