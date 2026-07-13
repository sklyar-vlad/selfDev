package auth

import (
	"context"
	"encoding/json"
	"net/http"

	model "github.com/sklyar-vlad/selfDev/internal/model/user"
	"go.uber.org/zap"
)

type AuthService interface {
	Auth(code, state string) (string, error)
	UserInfo(sub string) (string, error)
	FindOrCreate(context context.Context, userSub string) (model.User, error)
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

	userSub, err := h.service.UserInfo(accessToken)

	user, err := h.service.FindOrCreate(r.Context(), userSub)


	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   12 * 60 * 60,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   30 * 24 * 60 * 60,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err = json.NewEncoder(w).Encode(ToAuthResponse(&user)); err != nil {
		h.logger.Error("failed create response", zap.String("email", input.Email), zap.Error(err))
	}
}
