package utils

import "errors"

var (
	ErrEmptyAuthorizationToken = errors.New("error: Empty Authorization Token in header")
	ErrUnableToParseJWTToken   = errors.New("error: Unable to parse JWT token")
	ErrUnauthorized            = errors.New("error: Unauthorized")
	ErrInvalidAuthToken        = errors.New("error: Invalid JWT auth-token")
	ErrEmailAlreadyExists      = errors.New("email already exists")
	ErrAuthenticationFailure   = errors.New("error: Authentication failed")
	ErrIncorrectPassword       = errors.New("error: Password not correct")
	ErrorNotFound              = errors.New("error: Entity not found")
	ErrForbidden               = errors.New("error: Attempted action is not allowed")
	ErrUnableToCreateUser      = errors.New("error: Unable to create User")
	ErrUnableToFetchUser       = errors.New("error: Unable to fetch user")
	ErrUnableToFetchUserList   = errors.New("error: Unable to fetch user list")
	ErrUnableToUpdateUser      = errors.New("error: Unable to update user")
	ErrUnableToDeleteUser      = errors.New("error: Unable to delete user")

	// ErrResetExpired occurs when the reset hash exceeds the expiration
	ErrResetExpired = errors.New("error: Reset expired")
)
