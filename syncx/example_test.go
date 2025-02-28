package syncx_test

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/wwwangxc/wheel/syncx"
)

// ExampleWaitGroup example for syncx.WaitGroup
func ExampleWaitGroup() {
	var wg syncx.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		fmt.Println("  Goroutine[1] DONE")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(200 * time.Millisecond)
		fmt.Println("  Goroutine[2] DONE")
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println("Wait:")
	select {
	case <-ctx.Done():
		fmt.Println("  Context DONE")
	case <-wg.Wait():
		fmt.Println("  ALL DONE")
	}

	ctx1, cancel1 := context.WithTimeout(context.Background(), 111*time.Millisecond)
	defer cancel1()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		fmt.Println("  Goroutine[1] DONE")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(200 * time.Millisecond)
		fmt.Println("  Goroutine[2] DONE")
	}()

	fmt.Println("WaitOrDone:")
	if err := wg.WaitOrDone(ctx1); err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			fmt.Println("  Context Cancel")
		case errors.Is(err, context.DeadlineExceeded):
			fmt.Println("  Context Timeout")
		default:
			fmt.Println(err)
		}

		return
	}

	fmt.Println("ALL DONE")

	// Output:
	// Wait:
	//   Goroutine[1] DONE
	//   Goroutine[2] DONE
	//   ALL DONE
	// WaitOrDone:
	//   Goroutine[1] DONE
	//   Context Timeout
}
