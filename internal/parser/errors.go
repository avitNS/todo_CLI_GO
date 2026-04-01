package parser

import "errors"

var (
	ErrUnknownCommand = errors.New("unknown command")
	ErrMissingCommand = errors.New("command is missing")
)
