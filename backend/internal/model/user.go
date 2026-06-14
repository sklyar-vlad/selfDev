package model

import (
	"time"

	"uuid"
)

type User struct {
	user_id    uuid.UUID
	role       string
	username   string
	email      string
	password   string
	created_at time.Time
}
