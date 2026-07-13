package auth

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/sklyar-vlad/selfDev/internal/config"
	appErrors "github.com/sklyar-vlad/selfDev/internal/errors"
	auth "github.com/sklyar-vlad/selfDev/internal/integrations/casdoor"
	"github.com/sklyar-vlad/selfDev/internal/model/user"
	"go.uber.org/zap"
)

type UserService interface {
	GetUserByID(ctx context.Context, userSub string) (model.User, error)
	CreateUser(ctx context.Context, user model.User) (model.User, error)
}

type AuthAdapter interface {
	GetAccess(code, state string) (string, error)
	GetUserInfo(token string) (auth.AuthUser, error)
}

type AuthRepository interface {
	CreateSession(ctx context.Context, sessionID string, userID uuid.UUID) error
}

type Service struct {
	userService UserService
	authAdapter AuthAdapter
	repo        AuthRepository
	cfg         config.ConfigJWT
	logger      *zap.Logger
}

func NewService(
	userService UserService,
	authAdapter AuthAdapter,
	repo AuthRepository,
	configJwt config.ConfigJWT,
	logger *zap.Logger,
) *Service {
	return &Service{userService: userService, authAdapter: authAdapter, repo: repo, cfg: configJwt, logger: logger}
}

func (s *Service) Auth(code, state string) (string, error) {
	return s.authAdapter.GetAccess(code, state)
}

func (s *Service) GetUserInfo(token string) (auth.AuthUser, error) {
	return s.authAdapter.GetUserInfo(token)
}

func (s *Service) FindOrCreate(ctx context.Context, authUser auth.AuthUser) (model.User, error) {
	user, err := s.userService.GetUserByID(ctx, authUser.Sub)

	if err != nil {
		return model.User{}, err
	}

	if errors.Is(err, appErrors.ErrUserNotFound) {
		user, err = s.userService.CreateUser(ctx, model.NewUser(authUser.Sub, authUser.Name, authUser.Email))

		if err != nil {
			return model.User{}, err
		}

	}

	return user, nil
}

func (s *Service) CreateSession(ctx context.Context, userID uuid.UUID) (string, error) {
	sessionID := uuid.NewString()

	err := s.repo.CreateSession(ctx, sessionID, userID)

	if err != nil {
		return "", err
	}

	return sessionID, nil
}
