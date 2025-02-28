// Package coroutine helper
package coroutine

import (
	"github.com/wwwangxc/wheel"
)

// Go run the given function using another goroutine
func Go(fn func(), opts ...Option) {
	o := newOptions(opts...)
	wheel.DoIfNotNil(o.wg, func() { o.wg.Add(1) })
	wheel.DoIfNotNil(o.wgx, func() { o.wgx.Add(1) })

	go func() {
		defer func() {
			wheel.DoIfNotNil(o.wg, func() { o.wg.Done() })
			wheel.DoIfNotNil(o.wgx, func() { o.wgx.Done() })

			if err := recover(); err != nil {
				wheel.DoIfNotNil(o.logFn, func() { o.logFn(err) })
			}
		}()

		fn()
	}()
}
