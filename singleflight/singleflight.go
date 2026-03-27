// Package singleflight helper
package singleflight

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/singleflight"
)

var group singleflight.Group

// Do the given function with singleflight
//
// For details, see: https://pkg.go.dev/golang.org/x/sync/singleflight
func Do[T any](ctx context.Context,
	key string, fn func(context.Context) (T, error), opts ...Option) (T, error) {

	opt := newOptions(opts...)
	ch := group.DoChan(key, func() (any, error) {
		go func() {
			time.Sleep(opt.expire)
			group.Forget(key)
		}()

		return fn(ctx)
	})

	var zero T
	select {
	case <-ctx.Done():
		return zero, ctx.Err()
	case v := <-ch:
		if v.Err != nil {
			return zero, v.Err
		}

		ret, ok := v.Val.(T)
		if !ok {
			return zero, fmt.Errorf("the given function result type assert fail, expected:%T, actual:%T", zero, v.Val)
		}

		return ret, nil
	}
}
