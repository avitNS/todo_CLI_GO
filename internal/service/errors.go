package service

import "errors"

var (
	ErrInvalidTitle = errors.New("title is incorrect")
	ErrInvalidID    = errors.New("id is incorrect")
)
