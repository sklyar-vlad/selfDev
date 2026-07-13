package model

import (
	"github.com/google/uuid"
)

type User struct {
	UserId        uuid.UUID
	Sub			  string
	Username      string
	Email         string
}