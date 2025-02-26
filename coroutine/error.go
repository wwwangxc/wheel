package coroutine

import "errors"

var (
	// ErrTimeout indicates a timeout error
	ErrTimeout = errors.New("timeout")
)
