// Package coroutine helper
package coroutine

import (
	"fmt"

	"github.com/wwwangxc/wheel"
)

// Go run the given function using another goroutine
func Go(fn func() error, opts ...Option) {
	o := newOptions(opts...)
	wheel.DoIfNotNil(o.wg, func() { o.wg.Add(1) })

	go func() {
		defer wheel.DoIfNotNil(o.wg, func() { o.wg.Done() })
		runSafe(o.errCh, fn)
	}()
}

func runSafe(errCh chan<- error, fn func() error) {
	defer recoverx(errCh)
	if err := fn(); err != nil && errCh != nil {
		errCh <- err
	}
}

func recoverx(errCh chan<- error) {
	if err := recover(); err != nil && errCh != nil {
		errCh <- fmt.Errorf("%+v", err)
	}
}
