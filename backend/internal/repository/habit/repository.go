package habit

import (
	"context"
	"fmt"

	// "github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	model "github.com/sklyar-vlad/selfDev/internal/model/habit"
)

type repository struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewRepository(pool *pgxpool.Pool, logger *zap.Logger) *repository {
	return &repository{
		pool:   pool,
		logger: logger,
	}
}

// func (r *repository) GetAllHabits(ctx context.Context, userId uuid.UUID) (model.Habit, error) {
// 	return model.Habit{}, nil
// }

func (r *repository) CreateHabit(ctx context.Context, habit model.Habit) error {
	query := `
	INSERT INTO habits
	VALUES user_id = $1, habit_id = $2, name = $3, description = $4, category = $5
	`

	_, err := r.pool.Exec(ctx, query, habit.UserId, habit.HabitId, habit.Name, habit.Description, habit.Category)

	if err != nil {
		r.logger.Error("failed create habit in database", zap.Error(err))
		return fmt.Errorf("failed create habit in database: %v", err)
	}

}
