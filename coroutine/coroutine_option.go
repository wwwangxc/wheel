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

// WithErrChan set error channel, it will put error into the error chan if function return error or panics
func WithErrChan(ch chan<- error) Option {
	return func(o *options) {
		o.errCh = ch
	}
}

type options struct {
	wg    *sync.WaitGroup
	errCh chan<- error
}

func newOptions(opts ...Option) *options {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}

	return o
}
