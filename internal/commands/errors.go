package commands

import "errors"

var (
	ErrInvalidTitle = errors.New("title is invalid")
	ErrMissingTitle = errors.New("title is missing")
	ErrInvalidID    = errors.New("id is invalid")
	ErrMissingID    = errors.New("id is missing")
)
