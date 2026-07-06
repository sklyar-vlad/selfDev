package habit

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	model "github.com/sklyar-vlad/selfDev/internal/model/habit"
	userModel "github.com/sklyar-vlad/selfDev/internal/model/user"
)

type HabitRepository interface {
	GetAllHabits(ctx context.Context, userId uuid.UUID) (model.Habit, error)
}

type UserService interface {
	GetById(ctx context.Context, id uuid.UUID) (userModel.User, error)
}

type service struct {
	repo        HabitRepository
	userService UserService
	logger      *zap.Logger
}

func NewService(repo HabitRepository, userService UserService, logger *zap.Logger) *service {
	return &service{
		repo:        repo,
		userService: userService,
		logger:      logger,
	}
}

func (s *service) GetHabits(ctx context.Context, userId uuid.UUID) ([]model.Habit, error) {
	return nil, nil
}
