package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	appErrors "github.com/sklyar-vlad/selfDev/internal/errors"
	model "github.com/sklyar-vlad/selfDev/internal/model/user"
)

// TODO: Update(ctx context.Context, user model.User) (model.User, error)
// TODO: Delete(ctx context.Context, user model.User) error
type Repository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	GetByLogin(ctx context.Context, login string) (model.User, error)
	GetById(ctx context.Context, userId uuid.UUID) (model.User, error)
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
		return model.User{}, fmt.Errorf("failed generation password hash: %w", err)
	}

	user, err := model.NewUser(username, email, string(passwordHash))

	if errors.Is(err, appErrors.ErrInvalidEmail) {
		return model.User{}, appErrors.ErrInvalidEmail
	}

	if errors.Is(err, appErrors.ErrInvalidPassword) {
		return model.User{}, appErrors.ErrInvalidPassword
	}

	if err != nil {
		return model.User{}, fmt.Errorf("failed create user model: %w", err)
	}

	if err = s.repo.Create(ctx, &user); err != nil {
		return model.User{}, fmt.Errorf("failed insert user: %w", err)
	}

	s.logger.Info("success create user", zap.String("email", email))
	return user, nil
}

func (s *Service) UpdateUser(ctx context.Context, user *model.User) error {
	if err := s.repo.Update(ctx, user); err != nil {
		return err
	}

	s.logger.Info("success update user", zap.String("user_id", user.UserId.String()))
	return nil
}

func (s *Service) GetByLogin(ctx context.Context, username, email string) (model.User, error) {
	var login string
	if email == "" {
		login = username
	} else {
		login = email
	}

	user, err := s.repo.GetByLogin(ctx, login)

	if errors.Is(err, appErrors.ErrUserNotFound) {
		return model.User{}, appErrors.ErrUserNotFound
	}

	if err != nil {
		return model.User{}, fmt.Errorf("failed get user: %w", err)
	}

	s.logger.Info("success get user", zap.String("email", email))
	return user, nil
}

func (s *Service) GetById(ctx context.Context, userId uuid.UUID) (model.User, error) {
	user, err := s.repo.GetById(ctx, userId)

	if errors.Is(err, appErrors.ErrUserNotFound) {
		return model.User{}, appErrors.ErrUserNotFound
	}

	if err != nil {
		return model.User{}, fmt.Errorf("failed get user: %w", err)
	}

	s.logger.Info("success get user", zap.String("user_id", userId.String()))
	return user, nil
}
