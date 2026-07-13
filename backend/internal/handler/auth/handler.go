package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	auth "github.com/sklyar-vlad/selfDev/internal/integrations/casdoor"
	model "github.com/sklyar-vlad/selfDev/internal/model/user"
	"go.uber.org/zap"
)

type AuthService interface {
	Auth(code, state string) (string, error)
	GetUserInfo(sub string) (auth.AuthUser, error)
	FindOrCreate(ctx context.Context, user auth.AuthUser) (model.User, error)
	CreateSession(ctx context.Context, userId uuid.UUID) (string, error)
}

type handler struct {
	service AuthService
	logger  *zap.Logger
}

func NewHandler(service AuthService, logger *zap.Logger) *handler {
	return &handler{service: service, logger: logger}
}

func (h *handler) Callback(w http.ResponseWriter, r *http.Request) {
	var input AuthRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.logger.Error("failed decode request", zap.Error(err))
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	// if input.State != r.oauth2.State {
	// 	h.logger.Error("invalid oauth state")
	// 	http.Error(w, "invalid oauth state", http.StatusBadRequest)
	//	return
	// }

	accessToken, err := h.service.Auth(input.Code, input.State)
	if err != nil {
		h.logger.Error("failed auth", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	authUser, err := h.service.GetUserInfo(accessToken)
	user, err := h.service.FindOrCreate(r.Context(), authUser)
	sessionID, err := h.service.CreateSession(r.Context(), user.UserId)

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   30 * 24 * 3600,
	})
}
