package errorx

type options struct {
	reason string
}

type option func(*options)

func newOptions(opts ...option) *options {
	var o options
	for _, opt := range opts {
		opt(&o)
	}

	return &o
}

// WithReason set error reason
func WithReason(reason string) option {
	return func(o *options) {
		o.reason = reason
	}
}
