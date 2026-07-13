package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"github.com/sklyar-vlad/selfDev/internal/config"
	appErrors "github.com/sklyar-vlad/selfDev/internal/errors"
	userModel "github.com/sklyar-vlad/selfDev/internal/model/user"
)

type UserService interface {
	CreateUser(ctx context.Context, username, email, password string) (userModel.User, error)
	GetById(ctx context.Context, id uuid.UUID) (userModel.User, error)
}

type AuthAdapter interface {
	GetToken(code, state string) (string, error)
	GetUserSub(token string) (string, error)
}

type Service struct {
	userService  UserService
	authAdapter  AuthAdapter
	cfg          config.ConfigJWT
	logger       *zap.Logger
}

func NewService(
	userService UserService,
	authAdapter AuthAdapter,
	configJwt config.ConfigJWT,
	logger *zap.Logger,
) *Service {
	return &Service{userService: userService, authAdapter: authAdapter, cfg: configJwt, logger: logger}
}

func (s *Service) Auth(code, state string) (string, error) {
	return s.authAdapter.GetToken(code, state)
}

func (s *Service) GetUserInfo(token string) (string, error) {
	return s.authAdapter.GetUserSub(token)
}
