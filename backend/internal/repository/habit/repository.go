package habit

import (
	"context"
	"fmt"

	"github.com/google/uuid"
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


func (r *repository) GetAllHabits(ctx context.Context, userId uuid.UUID) ([]model.Habit, error) {
	query := `
	SELECT habit_id, name, description, is_good
	FROM habits
	WHERE user_id = $1
	ORDER BY created_at DESC
	`

	rows, err := r.pool.Query(ctx, query, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get habits: %w", err)
	}
	defer rows.Close()

	var habits []model.Habit

	for rows.Next() {
		var h model.Habit

		err := rows.Scan(
			&h.HabitId,
			&h.Name,
			&h.Description,
			&h.IsGood,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan habit: %w", err)
		}

		habits = append(habits, h)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return habits, nil
}

func (r *repository) CreateHabit(ctx context.Context, habit model.Habit) error {
	query := `
	INSERT INTO habits (user_id, habit_id, name, description, is_good)
	VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.pool.Exec(ctx, query, habit.UserId, habit.HabitId, habit.Name, habit.Description, habit.IsGood)
	if err != nil {
		r.logger.Error("failed create habit in database", zap.Error(err))
		return fmt.Errorf("failed create habit in database: %v", err)
	}

	return nil
}

func (r *repository) DeleteHabit(ctx context.Context, habitId uuid.UUID) error {
	query := `
	DELETE FROM habits
	WHERE habit_id = $1
	`

	_, err := r.pool.Exec(ctx, query, habitId)
	if err != nil {
		r.logger.Error("failed delete habit from database", zap.Error(err))
		return fmt.Errorf("failed delete habit from database: %v", err)
	}

	return nil
}