package errors

import (
	"errors"
)

var NoProfileExists = errors.New("no profile exists")

var ProfileNotSet = errors.New("user profile not set")

var TokenEmptyError = errors.New("token cannot be empty")

var RepoDetailsError = errors.New("username and repository name or Git remote URL missing")

func New(text string) error {
	return errors.New(text)
}

// FlagError is the kind of error raised in flag processing
type FlagError struct {
	Err error
}

func (fe FlagError) Error() string {
	return fe.Err.Error()
}

func (fe FlagError) Unwrap() error {
	return fe.Err
}
