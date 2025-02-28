package coroutine_test

import (
	"sync"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/wwwangxc/wheel/coroutine"
	"github.com/wwwangxc/wheel/syncx"
)

func TestGo(t *testing.T) {
	Convey("Given a function to be exected", t, func() {
		fn := func() {
			time.Sleep(100 * time.Millisecond)
		}

		Convey("Go with sync.WaitGroup", func() {
			var wg sync.WaitGroup
			coroutine.Go(fn, coroutine.WithWaitGroup(&wg))
			coroutine.Go(fn, coroutine.WithWaitGroup(&wg))

			Convey("Wait group should be waitting", func() {
				So(&wg, shouldBeWaitting)
			})

			Convey("Wait group should not be waitting", func() {
				time.Sleep(200 * time.Millisecond)
				So(&wg, shouldNotBeWaitting)
			})

		})

		Convey("Go with syncx.WaitGroup", func() {
			var wg syncx.WaitGroup
			coroutine.Go(fn, coroutine.WithWaitGroupX(&wg))
			coroutine.Go(fn, coroutine.WithWaitGroupX(&wg))

			Convey("Wait group should be waitting", func() {
				So(&wg, shouldBeWaitting)
			})

			Convey("Wait group should not be waitting", func() {
				time.Sleep(200 * time.Millisecond)
				So(&wg, shouldNotBeWaitting)
			})

		})
	})
}

func shouldBeWaitting(actual any, expected ...any) string {
	wait := func() {}

	switch g := actual.(type) {
	case *sync.WaitGroup:
		wait = func() { g.Wait() }
	case *syncx.WaitGroup:
		wait = func() { g.Wait() }
	case coroutine.Group:
		wait = func() { g.Wait() }
	default:
		return "The type of actual should be in [*sync.WaitGroup|*syncx.WaitGroup|coroutine.Group]"
	}

	done := make(chan struct{})
	go func() {
		wait()
		close(done)
	}()

	select {
	case <-done:
		return "Should be waitting"
	default:
	}

	return ""
}

func shouldNotBeWaitting(actual any, expected ...any) string {
	wait := func() {}

	switch g := actual.(type) {
	case *sync.WaitGroup:
		wait = func() { g.Wait() }
	case *syncx.WaitGroup:
		wait = func() { g.Wait() }
	case coroutine.Group:
		wait = func() { g.Wait() }
	default:
		return "The type of actual should be in [*sync.WaitGroup|*syncx.WaitGroup|coroutine.Group]"
	}

	done := make(chan struct{})
	go func() {
		wait()
		close(done)
	}()

	time.Sleep(time.Millisecond)
	select {
	case <-done:
		return ""
	default:
	}

	return "Should not be waitting"
}
