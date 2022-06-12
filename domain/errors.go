package domain

import "errors"

var (
	ErrInternal                = errors.New("internal error occured")
	ErrUnauthenticated         = errors.New("can't authenticate user")
	ErrAuthenticationFailed    = errors.New("user authentication failed")
	ErrAlreadyExists           = errors.New("resource already exists")
	ErrNotExists               = errors.New("resource not exists")
	ErrInvalidPaginationParams = errors.New("invalid pagination params")
)
