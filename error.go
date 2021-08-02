package gauth

import "errors"

var (
	// ErrSecretLengthLss secret length less than 6
	ErrSecretLengthLss = errors.New("secret length lss 6 error")
	// ErrParam param error
	ErrParam = errors.New("param error")
)
