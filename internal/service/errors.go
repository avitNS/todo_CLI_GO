package service

import "errors"

var (
	ErrMissingTitle = errors.New("Title is missing")
	ErrMissingID    = errors.New("ID is incorrect")
)
