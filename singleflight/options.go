package singleflight

import "time"

// Option of method do
type Option func(*options)

// WithExpire set expiration time
//
// Default `time.Second`
func WithExpire(expire time.Duration) Option {
	return func(o *options) {
		o.expire = expire
	}
}

type options struct {
	expire time.Duration
}

func newOptions(opts ...Option) *options {
	opt := defaultOptions()
	for _, v := range opts {
		v(opt)
	}

	return opt
}

func defaultOptions() *options {
	return &options{
		expire: time.Second,
	}
}
