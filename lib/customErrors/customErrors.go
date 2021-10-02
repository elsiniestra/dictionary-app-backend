package customErrors

import "github.com/pkg/errors"

var (
	IncorrectArgument     = errors.New("Incorrect argument provided")
	UnableFetchInstance   = errors.New("Unable fetch instance")
	FetchedInstanceIsNil  = errors.New("Fetched instance is nil")
	UnableCreateInstance  = errors.New("Unable create instance")
	UnableProcessInstance = errors.New("Unable process instance")
)
