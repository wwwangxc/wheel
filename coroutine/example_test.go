package coroutine_test

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/wwwangxc/wheel/coroutine"
)

func ExampleGo() {
	var wg sync.WaitGroup
	errCh := make(chan error)
	defer close(errCh)

	callback := func(err error) {
		fmt.Println(err)
		// do something...
	}

	coroutine.Go(
		func() {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("Goroutine[1] DONE")
		},
		coroutine.WithWaitGroup(&wg))

	coroutine.Go(
		func() {
			time.Sleep(150 * time.Millisecond)
			fmt.Println("Goroutine[2] DONE")
		},
		coroutine.WithWaitGroup(&wg))

	coroutine.Go(
		func() {
			time.Sleep(200 * time.Millisecond)
			fmt.Println("Goroutine[3] DONE")
			panic("panic message")
		},
		coroutine.WithWaitGroup(&wg),
		coroutine.WithCallbackWhenPanic(callback))

	wg.Wait()

	// Output:
	// Goroutine[1] DONE
	// Goroutine[2] DONE
	// Goroutine[3] DONE
	// panic message
}

func ExampleGroup() {
	fn := func(text string, sleep time.Duration, err error) func(context.Context) error {
		return func(ctx context.Context) error {
			time.Sleep(sleep)
			select {
			case <-ctx.Done():
				return nil
			default:
				fmt.Println(text)
			}

			return err
		}
	}
	f1 := fn("    - Goroutine[1] DONE", 100*time.Millisecond, nil)
	f2 := fn("    - Goroutine[2] DONE", 200*time.Millisecond, errors.New("error message"))
	f3 := fn("    - Goroutine[3] DONE", 300*time.Millisecond, nil)

	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second)
	defer cancel1()

	fmt.Println("Default")
	g := coroutine.NewGroup(ctx1)
	g.Go(f1)
	g.Go(f2)
	g.Go(f3)
	if err := g.Wait().Error(); err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second)
	defer cancel2()

	fmt.Println("Cancel On Error")
	g = coroutine.NewGroup(ctx2, coroutine.WithCancelOnError())
	g.Go(f1)
	g.Go(f2)
	g.Go(f3)
	if err := g.Wait().Error(); err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	ctx3, cancel3 := context.WithTimeout(context.Background(), time.Second)
	defer cancel3()

	fmt.Println("With Timeout")
	g = coroutine.NewGroup(ctx3, coroutine.WithTimeout(150*time.Millisecond))
	g.Go(f1)
	g.Go(f2)
	g.Go(f3)
	if err := g.Wait().Error(); err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	ctx4, cancel4 := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel4()

	fmt.Println("Context Timeout")
	g = coroutine.NewGroup(ctx4)
	g.Go(f1)
	g.Go(f2)
	g.Go(f3)
	if err := g.Wait().Error(); err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	// Output:
	// Default
	//     - Goroutine[1] DONE
	//     - Goroutine[2] DONE
	//     - Goroutine[3] DONE
	// 1 errors occurred:
	//     * error message
	//
	// Cancel On Error
	//     - Goroutine[1] DONE
	//     - Goroutine[2] DONE
	// 1 errors occurred:
	//     * error message
	//
	// With Timeout
	//     - Goroutine[1] DONE
	// coroutine group already timeout
	// 0 errors occurred:
	//
	// Context Timeout
	//     - Goroutine[1] DONE
	// coroutine group already timeout
	// 0 errors occurred:

}
