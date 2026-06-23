package user

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	model "github.com/sklyar-vlad/selfDev/internal/model/user"
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

func (r *repository) Create(ctx context.Context, user model.User) (model.User, error) {
	query := `
	INSERT INTO users (role, username, email, password)
	VALUES ($1, $2, $3,	$4)
	`

	_, err := r.pool.Exec(ctx, query, user.Role, user.Username, user.Email, user.Password)
	if err != nil {
		r.logger.Error("failed insert user in database", zap.Error(err))
		return model.User{}, fmt.Errorf("failed insert user in database: %v", err)
	}

	r.logger.Info("success insert user in database", zap.String("email", user.Email))

	return user, nil
}

// func (r *repository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
// 	query := `
// 	SELECT user_id, role, username, email, password
// 	FROM users
// 	WHERE email = $1
// 	`

// 	var user model.User

// 	err := r.pool.QueryRow(ctx, query, email).Scan(
// 		&user.UserId,
// 		&user.Role,
// 		&user.Username,
// 		&user.Email,
// 		&user.Password,
// 	)

// 	if errors.Is(err, pgx.ErrNoRows) {
// 		r.logger.Error("user not found", zap.Error(customErrors.ErrUserNotFound))
// 		return model.User{}, customErrors.ErrUserNotFound
// 	}

// 	if err != nil {
// 		r.logger.Error("failed get user by email", zap.Error(err))
// 		return model.User{}, err
// 	}

// 	r.logger.Info("success select user by email", zap.String("email", user.Email))
// 	return user, nil
// }
