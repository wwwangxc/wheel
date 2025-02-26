package coroutine_test

import (
	"context"
	"errors"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/wwwangxc/wheel/coroutine"
)

func TestGroup(t *testing.T) {
	Convey("Given some global arguments", t, func() {
		ctx := context.Background()
		errMessage := "this is an error message"
		fn := func(timeout time.Duration) func(context.Context) error {
			return func(_ context.Context) error {
				time.Sleep(timeout)
				return errors.New(errMessage)
			}
		}

		Convey("New group with CancelOnError option", func() {
			g := coroutine.NewGroup(ctx,
				coroutine.WithCancelOnError(),
				coroutine.WithConcurrencyLevel(1))

			Convey("3 functions are running with 1 goroutine", func() {
				g.Go(fn(100 * time.Millisecond))
				g.Go(fn(100 * time.Millisecond))
				g.Go(fn(100 * time.Millisecond))

				Convey("Group should be waitting", func() {
					So(g, shouldBeWaitting)
				})

				Convey("Group should not be waitting", func() {
					err := g.Wait()
					So(g, shouldNotBeWaitting)

					Convey("Error number should be 1", func() {
						So(len(err.Errors()), ShouldEqual, 1)
					})
				})
			})
		})

		Convey("New group has no CancelOnError option", func() {
			g := coroutine.NewGroup(ctx, coroutine.WithConcurrencyLevel(3))

			Convey("3 functions are running with 3 goroutine", func() {
				g.Go(fn(1 * time.Millisecond))
				g.Go(fn(11 * time.Millisecond))
				g.Go(fn(111 * time.Millisecond))

				Convey("Group should be waitting", func() {
					So(g, shouldBeWaitting)
				})

				Convey("Group should not be waitting", func() {
					err := g.Wait()
					So(g, shouldNotBeWaitting)

					Convey("Error number should be 3", func() {
						So(len(err.Errors()), ShouldEqual, 3)
					})
				})
			})
		})

		Convey("New group with timeout 10ms", func() {
			g := coroutine.NewGroup(ctx,
				coroutine.WithTimeout(10*time.Millisecond),
				coroutine.WithConcurrencyLevel(3))

			Convey("One function running", func() {
				g.Go(fn(100 * time.Millisecond))

				Convey("Group should be waitting", func() {
					So(g, shouldBeWaitting)
				})

				Convey("Group already timeout", func() {
					err := g.Wait()
					So(g, shouldNotBeWaitting)

					Convey("Error should be timeout", func() {
						So(errors.Is(err.Error(), coroutine.ErrTimeout), ShouldBeTrue)
						//So(err.Error(), ShouldBeError, coroutine.ErrTimeout)
					})
				})
			})
		})
	})
}
