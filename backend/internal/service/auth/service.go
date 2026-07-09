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
	authModel "github.com/sklyar-vlad/selfDev/internal/model/auth"
	userModel "github.com/sklyar-vlad/selfDev/internal/model/user"
)

type UserService interface {
	CreateUser(ctx context.Context, username, email, password string) (userModel.User, error)
	UpdateUser(ctx context.Context, user *userModel.User) error
	GetByLogin(ctx context.Context, username, password string) (userModel.User, error)
	GetById(ctx context.Context, id uuid.UUID) (userModel.User, error)
}

type EmailAdapter interface {
	SendEmailVerification(email, token string) error
}

type Repository interface {
	CreateRefreshToken(ctx context.Context, Tokens *authModel.Tokens) error
	GetRefreshToken(ctx context.Context, userId uuid.UUID) (authModel.Tokens, error)
	DeleteRefreshToken(ctx context.Context, userId uuid.UUID) error
	SaveTokenVerify(ctx context.Context, token, userId string) error
	ConsumeToken(ctx context.Context, token string) (string, error)
}

type Service struct {
	repo         Repository
	userService  UserService
	emailAdapter EmailAdapter
	cfg          config.ConfigJWT
	logger       *zap.Logger
}

func NewService(
	repo Repository,
	userService UserService,
	emailAdapter EmailAdapter,
	configJwt config.ConfigJWT,
	logger *zap.Logger,
) *Service {
	return &Service{repo: repo, userService: userService, emailAdapter: emailAdapter, cfg: configJwt, logger: logger}
}

func (s *Service) Login(ctx context.Context, username, email, password string) (string, string, error) {
	user, err := s.userService.GetByLogin(ctx, username, email)

	if errors.Is(err, appErrors.ErrUserNotFound) {
		return "", "", appErrors.ErrUserNotFound
	}

	if err != nil {
		return "", "", err
	}

	if !user.EmailVerified {
		return "", "", appErrors.ErrEmailNotVerified
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", appErrors.ErrInvalidPassword
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, authModel.NewRefreshClaims(user.UserId)).
		SignedString([]byte(s.cfg.Secret))
	if err != nil {
		return "", "", fmt.Errorf("failed hash generation: %w", err)
	}

	refreshTokenHash := sha256.Sum256([]byte(refreshToken))

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, authModel.NewAccessClaims(user.UserId)).
		SignedString([]byte(s.cfg.Secret))
	if err != nil {
		return "", "", fmt.Errorf("failed signed token: %w", err)
	}

	var tokens authModel.Tokens
	tokens.AccessToken = accessToken
	tokens.RefreshToken = hex.EncodeToString(refreshTokenHash[:])
	tokens.ExpiresAt = time.Now().AddDate(0, 1, 0)
	tokens.UserId = user.UserId

	err = s.repo.CreateRefreshToken(ctx, &tokens)
	if err != nil {
		return "", "", fmt.Errorf("failed create refresh token: %w", err)
	}

	s.logger.Info("success login", zap.String("email", user.Email))
	return refreshToken, accessToken, nil
}

func (s *Service) Logout(ctx context.Context, refreshToken string) error {
	token, err := jwt.ParseWithClaims(
		refreshToken,
		&authModel.Claims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(s.cfg.Secret), nil
		},
	)
	if err != nil {
		return fmt.Errorf("invalid validate jwt token: %w", err)
	}

	claims, ok := token.Claims.(*authModel.Claims)
	if !ok || !token.Valid {
		return errors.New("invalid token")
	}

	err = s.repo.DeleteRefreshToken(ctx, claims.UserId)
	if err != nil {
		return fmt.Errorf("failed delete refresh token: %w", err)
	}

	s.logger.Info("success logout", zap.String("user_id", claims.UserId.String()))
	return nil
}

func (s *Service) Refresh(ctx context.Context, refreshToken string) (string, error) {
	token, err := jwt.ParseWithClaims(
		refreshToken,
		&authModel.Claims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(s.cfg.Secret), nil
		},
	)
	if err != nil {
		return "", fmt.Errorf("invalid validate jwt token: %w", err)
	}

	claims, ok := token.Claims.(*authModel.Claims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	refreshTokenInDB, err := s.repo.GetRefreshToken(ctx, claims.UserId)
	if err != nil {
		return "", err
	}

	refreshTokenHash := sha256.Sum256([]byte(refreshToken))

	if hex.EncodeToString(refreshTokenHash[:]) != refreshTokenInDB.RefreshToken {
		return "", fmt.Errorf("invalid refresh token: %w", err)
	}

	if time.Now().After(refreshTokenInDB.ExpiresAt) {
		return "", appErrors.ErrTokenWasExpired
	}

	newAccessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, authModel.NewAccessClaims(claims.UserId)).
		SignedString([]byte(s.cfg.Secret))
	if err != nil {
		return "", fmt.Errorf("failed signed token: %w", err)
	}

	s.logger.Info("success refresh", zap.String("user_id", claims.UserId.String()))
	return newAccessToken, nil
}

func (s *Service) Register(ctx context.Context, username, email, password string) error {
	user, err := s.userService.CreateUser(ctx, username, email, password)

	if errors.Is(err, appErrors.ErrEmailAlreadyExists) {
		return appErrors.ErrEmailAlreadyExists
	}

	if errors.Is(err, appErrors.ErrUsernameAlreadyExists) {
		return appErrors.ErrUsernameAlreadyExists
	}

	if err != nil {
		return fmt.Errorf("failed create user: %w", err)
	}

	token, err := authModel.NewTokenVerify()
	if err != nil {
		return fmt.Errorf("failed create verify token: %w", err)
	}

	err = s.repo.SaveTokenVerify(ctx, token.TokenVer, user.UserId.String())
	if err != nil {
		return fmt.Errorf("failed save verify token: %w", err)
	}

	go func() {
		err := s.emailAdapter.SendEmailVerification(email, token.TokenVer)
		if err != nil {
			s.logger.Error("failed send verification message", zap.Error(err))
		}
	}()

	s.logger.Info("success register", zap.String("email", user.Email))
	return nil
}

func (s *Service) ConfirmEmail(ctx context.Context, token string) error {
	userId, err := s.repo.ConsumeToken(ctx, token)
	if err != nil {
		return fmt.Errorf("failed consume token: %w", err)
	}

	userIdUUID, err := uuid.Parse(userId)
	if err != nil {
		return fmt.Errorf("invalid uuid: %w", err)
	}

	userEmailVerified, err := s.userService.GetById(ctx, userIdUUID)
	if err != nil {
		return fmt.Errorf("failed get user: %v", err)
	}

	userEmailVerified.EmailVerified = true

	if err = s.userService.UpdateUser(ctx, &userEmailVerified); err != nil {
		return fmt.Errorf("failed verify email: %v", err)
	}

	s.logger.Info("success confirm email", zap.String("email", userEmailVerified.Email))
	return nil
}

func (s *Service) GetCurrentUser(ctx context.Context, accessToken string) (userModel.User, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&authModel.Claims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(s.cfg.Secret), nil
		},
	)
	if err != nil {
		return userModel.User{}, fmt.Errorf("invalid access token: %w", err)
	}

	claims, ok := token.Claims.(*authModel.Claims)
	if !ok || !token.Valid {
		return userModel.User{}, errors.New("invalid token")
	}

	user, err := s.userService.GetById(ctx, claims.UserId)
	if err != nil {
		return userModel.User{}, fmt.Errorf("failed get current user: %w", err)
	}

	s.logger.Info("success get current user", zap.String("email", user.Email))
	return user, nil
}
