package repo_errors

import "errors"

var (
	ErrCommandNotFound = errors.New("not found value")
	ErrCreatingCommand = errors.New("creating command")
)
