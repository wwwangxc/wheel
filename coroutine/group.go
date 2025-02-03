package coroutine

import (
	"context"
	"sync"
)

// GoFunc is the function executed in each coroutine
type GoFunc func(ctx context.Context) error

// Group of coroutine task
type Group interface {
	// Go run the given function using another goroutine, recovers and return error if function panics.
	Go(fn GoFunc)
	// Wait until all given function have finished executing
	Wait() *GroupError
}

// NewGroup create the coroutine group
func NewGroup(ctx context.Context, opts ...GroupOption) Group {
	o := newGroupOptions(opts...)
	subCtx, cancel := context.WithTimeout(ctx, o.timeout)
	g := &groupImpl{
		ctx:           subCtx,
		ctxCancel:     sync.OnceFunc(cancel),
		cancelOnError: o.cancelOnError,
		err:           newGroupError(),
		errCh:         make(chan error, 1),
		rateLimit:     make(chan struct{}, o.concurrencyLevel),
	}

	for i := 0; i < o.concurrencyLevel; i++ {
		g.rateLimit <- struct{}{}
	}

	g.watchError()
	return g
}

type groupImpl struct {
	ctx           context.Context
	ctxCancel     context.CancelFunc
	cancelOnError bool
	err           *GroupError
	errCh         chan error
	wg            sync.WaitGroup
	rateLimit     chan struct{}
}

// Go run the given function using another goroutine, recovers and return error if function panics.
func (s *groupImpl) Go(fn GoFunc) {
	if s.alreadyDone() {
		return
	}

	select {
	case <-s.ctx.Done():
		return
	case <-s.rateLimit:
	}

	Go(
		func() error {
			defer func() {
				s.rateLimit <- struct{}{}
			}()

			return fn(s.ctx)
		},
		WithWaitGroup(&s.wg),
		WithErrChan(s.errCh),
	)
}

// Wait until all given function have finished executing
func (s *groupImpl) Wait() *GroupError {
	done := make(chan struct{})

	wait := func() error {
		s.wg.Wait()
		close(done)
		return nil
	}
	Go(wait)

	select {
	case <-s.ctx.Done():
		return s.err
	case <-done:
		return s.err
	}
}

func (s *groupImpl) watchError() {
	if s.alreadyDone() {
		return
	}

	Go(func() error {
		for {
			select {
			case <-s.ctx.Done():
				return nil
			case err, ok := <-s.errCh:
				if !ok {
					return nil
				}

				s.err.append(err)
				if s.cancelOnError {
					s.ctxCancel()
				}
			}
		}
	})
}

func (s *groupImpl) alreadyDone() bool {
	if s == nil {
		return true
	}

	select {
	case <-s.ctx.Done():
		return true
	default:
	}

	return false
}
