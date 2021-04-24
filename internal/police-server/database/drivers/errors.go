package drivers

import "errors"

var (
	ErrEmptyStruct         = errors.New("empty structure")
	ErrInvalidConfigStruct = errors.New("invalid config structure")
)
