package coroutine

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/wwwangxc/wheel"
)

// GroupError is a collection of all errors returned by the given function
type GroupError struct {
	errs      []error
	rw        sync.RWMutex
	isTimeout bool
}

func newGroupError() *GroupError {
	return &GroupError{
		errs: []error{},
	}
}

// Error return merged error of all errors returned by the given function
func (s *GroupError) Error() error {
	if s == nil || (!s.isTimeout && len(s.errs) == 0) {
		return nil
	}

	s.rw.RLock()
	defer s.rw.RUnlock()

	num := 0
	var buf bytes.Buffer
	for _, err := range s.errs {
		num++
		fmt.Fprintf(&buf, "\n    * %+v", err)
	}

	err := fmt.Errorf("%d errors occurred:%s", num, buf.String())
	if s.isTimeout {
		err = fmt.Errorf("coroutine group already %w\n%w", ErrTimeout, err)
	}

	return err
}

// Errors return a collection of all errors returned by the given function
func (s *GroupError) Errors() []error {
	if s == nil || (!s.isTimeout && len(s.errs) == 0) {
		return nil
	}

	s.rw.RLock()
	defer s.rw.RUnlock()

	errs := make([]error, len(s.errs))
	copy(errs, s.errs)
	if s.isTimeout {
		errs = append([]error{ErrTimeout}, errs...)
	}

	return errs
}

func (s *GroupError) append(err error) {
	if s == nil || err == nil {
		return
	}

	s.rw.Lock()
	defer s.rw.Unlock()

	s.errs = append(s.errs, err)
}

func (s *GroupError) alreadyTimeout() {
	wheel.DoIfNotNil(s, func() {
		s.isTimeout = true
	})
}
