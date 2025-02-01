package coroutine

import (
	"runtime"
	"time"
)

// GroupOption is the configuration option of the group
type GroupOption func(o *groupOptions)

// WithConcurrencyLevel set the concurrency level for the group. Default cpuNum*10
func WithConcurrencyLevel(level int) GroupOption {
	return func(o *groupOptions) {
		o.concurrencyLevel = level
	}
}

// WithCancelOnError exit immediately when any coroutine returns an error. Default false
func WithCancelOnError() GroupOption {
	return func(o *groupOptions) {
		o.cancelOnError = true
	}
}

// WithTimeout set the global timeout for the group. Default 3s
func WithTimeout(timeout time.Duration) GroupOption {
	return func(o *groupOptions) {
		o.timeout = timeout
	}
}

type groupOptions struct {
	concurrencyLevel int
	cancelOnError    bool
	timeout          time.Duration
}

func newGroupOptions(opts ...GroupOption) *groupOptions {
	o := defaultGroupOptions()
	for _, opt := range opts {
		opt(o)
	}

	return o
}

func defaultGroupOptions() *groupOptions {
	return &groupOptions{
		concurrencyLevel: runtime.NumCPU() * 10,
		cancelOnError:    false,
		timeout:          time.Second * 3,
	}
}
