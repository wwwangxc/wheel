package coroutine

import (
	"bytes"
	"fmt"
	"sync"
)

// GroupError is a collection of all errors returned by the given function
type GroupError struct {
	errs []error
	rw   sync.RWMutex
}

func newGroupError() *GroupError {
	return &GroupError{
		errs: []error{},
	}
}

// Error return merged error of all errors returned by the given function
func (s *GroupError) Error() error {
	if s == nil || len(s.errs) == 0 {
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

	return fmt.Errorf("%d errors occurred:%s", num, buf.String())
}

// Errors return a collection of all errors returned by the given function
func (s *GroupError) Errors() []error {
	if s == nil || len(s.errs) == 0 {
		return nil
	}

	s.rw.RLock()
	defer s.rw.RUnlock()

	return s.errs
}

func (s *GroupError) append(err error) {
	if s == nil || err == nil {
		return
	}

	s.rw.Lock()
	defer s.rw.Unlock()

	s.errs = append(s.errs, err)
}
