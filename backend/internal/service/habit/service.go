package habit

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.uber.org/zap"

	model "github.com/sklyar-vlad/selfDev/internal/model/habit"
	userModel "github.com/sklyar-vlad/selfDev/internal/model/user"
)

type HabitRepository interface {
	GetAllHabits(ctx context.Context, userId uuid.UUID) ([]model.Habit, error)
	CreateHabit(ctx context.Context, habit model.Habit) error
	DeleteHabit(ctx context.Context, habitId uuid.UUID) error
	ConfirmHabit(ctx context.Context, habitId uuid.UUID) error
	CancelHabit(ctx context.Context, habitId uuid.UUID) error
	GetHabitConfirmDates(ctx context.Context, habitId uuid.UUID) ([]model.Date, error)
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
	habits, err := s.repo.GetAllHabits(ctx, userId)
	if err != nil {
		s.logger.Error("failed get habits by user_id", zap.Error(err))
		return []model.Habit{}, fmt.Errorf("failed get habits by user_id: %v", err)
	}

	return habits, nil
}

func (s *service) CreateHabit(
	ctx context.Context,
	userId uuid.UUID,
	name, description string,
	isGood bool,
) (model.Habit, error) {
	habit, err := model.NewHabit(userId, name, description, isGood)
	if err != nil {
		s.logger.Error("failed create model habit", zap.Error(err))
		return model.Habit{}, fmt.Errorf("failed create model habit: %v", err)
	}

	if err = s.repo.CreateHabit(ctx, habit); err != nil {
		s.logger.Error("failed insert habit in database", zap.Error(err))
		return model.Habit{}, fmt.Errorf("failed insert habit in database: %v", err)
	}

	return habit, nil
}

func (s *service) DeleteHabit(ctx context.Context, habitId uuid.UUID) error {
	if err := s.repo.DeleteHabit(ctx, habitId); err != nil {
		s.logger.Error("failed delete habit from database", zap.Error(err))
		return fmt.Errorf("failed delete habit from database: %v", err)
	}

	return nil
}

func (s *service) ConfirmHabit(ctx context.Context, habitId uuid.UUID) error {
	if err := s.repo.ConfirmHabit(ctx, habitId); err != nil {
		s.logger.Error("failed confirm date of habit", zap.Error(err))
		return fmt.Errorf("failed confirm date of habit: %v", err)
	}

	return nil
}

func (s *service) CancelHabit(ctx context.Context, habitId uuid.UUID) error {
	if err := s.repo.CancelHabit(ctx, habitId); err != nil {
		s.logger.Error("failed cancel date of habit", zap.Error(err))
		return fmt.Errorf("failed cancel date of habit: %v", err)
	}

	return nil
}

func (s *service) GetHabitConfirmDates(ctx context.Context, habitId uuid.UUID) ([]model.Date, error) {
	dates, err := s.repo.GetHabitConfirmDates(ctx, habitId)
	if err != nil {
		s.logger.Error("failed cancel date of habit", zap.Error(err))
		return []model.Date{}, fmt.Errorf("failed cancel date of habit: %v", err)
	}

	return dates, nil
}
