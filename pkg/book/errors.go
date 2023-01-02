package book

import "errors"

var (
	ErrDatabaseRequired = errors.New("database is required")
	ErrRouterRequired   = errors.New("router is required")
)
