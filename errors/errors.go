package errors

import (
	"errors"
)

var (
	ErrGreaterThanLimit = errors.New("number is greater than upper limit allowed")
	ErrSmallerThanLimit = errors.New("number is smaller than the lower limit allowed")
	ErrBadFormatNumber  = errors.New("wrong number format")
)
