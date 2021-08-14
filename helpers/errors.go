package helpers

import "errors"

var (
	ErrInvalidEmail          = errors.New("invalid email")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrEmptyPassword         = errors.New("password can't be empty")
	ErrEmptyUsername         = errors.New("username can't be empty")
	ErrInvalidAuthToken      = errors.New("invalid auth-token")
	ErrInvalidCredentials    = errors.New("invalid credentials")
	ErrUnauthorized          = errors.New("Unauthorized")
)
