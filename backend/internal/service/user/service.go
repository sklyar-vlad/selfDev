package user

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	appErrors "github.com/sklyar-vlad/selfDev/internal/errors"
	model "github.com/sklyar-vlad/selfDev/internal/model/user"
)

type Repository interface {
	Create(ctx context.Context, user model.User) (model.User, error)
	// GetByLogin(ctx context.Context, user model.User) (model.User, error)
	// Update(ctx context.Context, user model.User) (model.User, error)
	// Delete(ctx context.Context, user model.User) error
}

type Service struct {
	repo   Repository
	logger *zap.Logger
}

func NewService(repo Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) CreateUser(ctx context.Context, username, email, password string) (model.User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error("failed password hash generation", zap.String("email", email), zap.Error(err))
		return model.User{}, fmt.Errorf("failed password hash generation: %v", err)
	}

	user, err := model.NewUser(username, email, string(passwordHash))

	if errors.Is(err, appErrors.ErrInvalidEmail) {
		s.logger.Error("invalid email", zap.Error(err))
		return model.User{}, appErrors.ErrInvalidEmail
	}

	if errors.Is(err, appErrors.ErrInvalidPassword) {
		s.logger.Error("invalid password", zap.Error(err))
		return model.User{}, appErrors.ErrInvalidPassword
	}

	if err != nil {
		s.logger.Error("failed create user model", zap.String("email", email), zap.Error(err))
		return model.User{}, fmt.Errorf("failed create user model: %v", err)
	}

	return s.repo.Create(ctx, user)
}
