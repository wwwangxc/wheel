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

// WithCallbackWhenPanic set callback func, it will be called when panic
func WithCallbackWhenPanic(fn func(error)) Option {
	return func(o *options) {
		o.callbackWhenPanic = fn
	}
}

type options struct {
	wg                *sync.WaitGroup
	wgx               *syncx.WaitGroup
	callbackWhenPanic func(error)
}

func newOptions(opts ...Option) *options {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}

	return o
}
