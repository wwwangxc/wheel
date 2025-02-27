package coroutine

import "sync"

// Option of coroutine
type Option func(o *options)

// WithWaitGroup set sync.WaitGroup
func WithWaitGroup(wg *sync.WaitGroup) Option {
	return func(o *options) {
		o.wg = wg
	}
}

// WithLogWhenPanic set log func, it will be called when panic
func WithLogWhenPanic(logFn func(v ...any)) Option {
	return func(o *options) {
		o.logFn = logFn
	}
}

type options struct {
	wg    *sync.WaitGroup
	logFn func(v ...any)
}

func newOptions(opts ...Option) *options {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}

	return o
}
