package syncx_test

import (
	"context"
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
		fmt.Println("Groutine[1] DONE")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Groutine[2] DONE")
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println("Timeout")
	case <-wg.Wait():
		fmt.Println("ALL DONE")
	}

	// Output:
	// Groutine[1] DONE
	// Groutine[2] DONE
	// ALL DONE
}
