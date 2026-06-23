package dto

import (
	"github.com/google/uuid"

	model "github.com/sklyar-vlad/selfDev/internal/model/user"
)

type UserResponse struct {
	UserId   uuid.UUID `json:"user_id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func ToUserResponse(p model.User) UserResponse {
	resp := UserResponse{
		UserId:   p.UserId,
		Email:    p.Email,
		Password: p.Password,
	}

	return resp
}
