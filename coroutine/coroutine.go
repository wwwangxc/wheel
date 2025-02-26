// Package coroutine helper
package coroutine

import (
	"github.com/wwwangxc/wheel"
)

// Go run the given function using another goroutine
func Go(fn func() error, opts ...Option) {
	o := newOptions(opts...)
	wheel.DoIfNotNil(o.wg, func() { o.wg.Add(1) })

	go func() {
		defer func() {
			wheel.DoIfNotNil(o.wg, func() { o.wg.Done() })
			if err := recover(); err != nil && o.errCh != nil {
				o.errCh <- err.(error)
			}
		}()
		wheel.MustBeNil(fn())
	}()
}
