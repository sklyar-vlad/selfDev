package errors

import "errors"

var (
	UserNotFound  = errors.New("User not found")
	EmailNotFound = errors.New("Email not found")
)
