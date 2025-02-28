package coroutine

import (
	"sync"

	"github.com/wwwangxc/wheel"
	"github.com/wwwangxc/wheel/syncx"
)

// Option of coroutine
type Option func(o *options)

// WithWaitGroup set sync.WaitGroup
func WithWaitGroup(wg *sync.WaitGroup) Option {
	return func(o *options) {
		wheel.DoIfNotNil(wg, func() {
			o.wg = wg
		})
	}
}

// WithWaitGroupX set syncx.WaitGroup
func WithWaitGroupX(wg *syncx.WaitGroup) Option {
	return func(o *options) {
		wheel.DoIfNotNil(wg, func() {
			o.wgx = wg
		})
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
	wgx   *syncx.WaitGroup
	logFn func(v ...any)
}

func newOptions(opts ...Option) *options {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}

	return o
}
