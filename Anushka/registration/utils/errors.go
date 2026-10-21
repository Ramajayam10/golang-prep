package utils

import "errors"

// Common Sentinel Errors
var (
	ErrIneligibleUser     = errors.New("user is not eligible to register: must be at least 18 with valid email")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUserNotFound       = errors.New("user not found")
	ErrTokenGeneration    = errors.New("failed to generate authentication token")
)
