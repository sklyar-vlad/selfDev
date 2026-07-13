package auth

import (
	"github.com/google/uuid"
	model "github.com/sklyar-vlad/selfDev/internal/model/user"
)

type AuthRequest struct {
	Code  string	  `json:"code"`
	State string	  `json:"state"`
}

type AuthResponse struct {
	UserId   uuid.UUID `json:"user_id"`
	Role     string    `json:"role"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

func ToAuthResponse(p *model.User) AuthResponse {
	resp := AuthResponse{
		UserId:   p.UserId,
		Username: p.Username,
		Email:    p.Email,
	}

	return resp
}
