package customerrors

import "github.com/pkg/errors"

var (
	ErrIncorrectArgument     = errors.New("Incorrect argument provided")
	ErrUnableFetchInstance   = errors.New("Unable fetch instance")
	ErrFetchedInstanceIsNil  = errors.New("Fetched instance is nil")
	ErrUnableCreateInstance  = errors.New("Unable create instance")
	ErrUnableProcessInstance = errors.New("Unable process instance")
)
