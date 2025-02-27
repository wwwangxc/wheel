package coroutine

import (
	"context"
	"errors"
	"sync"

	"github.com/wwwangxc/wheel"
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
		errQ:          make(chan error, o.concurrencyLevel),
		fnQ:           make(chan GoFunc, 100),
		rateLimit:     make(chan struct{}, o.concurrencyLevel),
	}

	for i := 0; i < o.concurrencyLevel; i++ {
		g.rateLimit <- struct{}{}
	}

	g.watch()
	return g
}

type groupImpl struct {
	ctx           context.Context
	ctxCancel     context.CancelFunc
	cancelOnError bool
	err           *GroupError
	errQ          chan error
	fnQ           chan GoFunc
	fnWG          sync.WaitGroup
	wg            sync.WaitGroup
	rateLimit     chan struct{}
}

// Go run the given function using another goroutine, recovers and return error if function panics.
func (s *groupImpl) Go(fn GoFunc) {
	if s.alreadyDone() {
		return
	}

	s.fnQ <- fn
	s.fnWG.Add(1)
}

// Wait until all given function have finished executing
func (s *groupImpl) Wait() *GroupError {
	allFnDone := make(chan struct{})

	Go(func() {
		s.fnWG.Wait()
		s.wg.Wait()
		close(allFnDone)
	})

	select {
	case <-s.ctx.Done():
		if errors.Is(s.ctx.Err(), context.DeadlineExceeded) {
			s.err.alreadyTimeout()
		}
	case <-allFnDone:
	}

	return s.err
}

func (s *groupImpl) watch() {
	wheel.DoIfNotNil(s, func() {
		s.watchFn()
		s.watchError()
	})
}

func (s *groupImpl) watchFn() {
	if s.alreadyDone() {
		return
	}

	Go(func() {
		for {
			var fn GoFunc
			var ok bool
			select {
			case <-s.ctx.Done():
				return
			case fn, ok = <-s.fnQ:
				if !ok {
					return
				}
				s.fnWG.Done()
			}

			select {
			case <-s.ctx.Done():
				return
			case _, ok = <-s.rateLimit:
				if !ok {
					return
				}
			}

			Go(
				func() {
					defer func() {
						s.rateLimit <- struct{}{}
					}()

					if err := fn(s.ctx); err != nil {
						s.errQ <- err
					}
				},
				WithWaitGroup(&s.wg))
		}
	})
}

func (s *groupImpl) watchError() {
	if s.alreadyDone() {
		return
	}

	Go(func() {
		for err := range s.errQ {
			s.err.append(err)
			if s.cancelOnError {
				s.ctxCancel()
				return
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
