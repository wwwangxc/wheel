package singleflight_test

import (
	"context"
	"fmt"
	"time"

	"github.com/wwwangxc/wheel/singleflight"
)

func Example() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ret, err := singleflight.Do(ctx, "function_key",
		func(ctx context.Context) (string, error) {
			// dosomething...
			time.Sleep(2 * time.Second)
			return "Successfully", nil
		}) // the function for key `function_key` will expire in 1 second

	switch {
	case err != nil:
		fmt.Println(err.Error())
	default:
		fmt.Println(ret)
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ret, err = singleflight.Do(ctx, "function_key",
		func(ctx context.Context) (string, error) {
			// dosomething...
			return "Successfully", nil
		}, singleflight.WithExpire(time.Minute)) // the function for key `function_key` will expire in 1 minute

	switch {
	case err != nil:
		fmt.Println(err.Error())
	default:
		fmt.Println(ret)
	}

	// Output:
	// context deadline exceeded
	// Successfully
}
