package coroutine_test

import (
	"errors"
	"sync"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/wwwangxc/wheel/coroutine"
)

func TestGo(t *testing.T) {
	Convey("Given a function return error to be exected", t, func() {
		errMessage := "this is an error message"
		fn := func() error {
			time.Sleep(100 * time.Millisecond)
			return errors.New(errMessage)
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

		Convey("Go with error channel", func() {
			ch := make(chan error, 1)
			coroutine.Go(fn, coroutine.WithErrChan(ch))
			time.Sleep(100 * time.Millisecond)
			So(len(ch), ShouldEqual, 1)
			err := <-ch
			So(err.Error(), ShouldEqual, errMessage)
		})
	})
}

func shouldBeWaitting(actual any, expected ...any) string {
	wait := func() {}

	switch g := actual.(type) {
	case *sync.WaitGroup:
		wait = func() { g.Wait() }
	case coroutine.Group:
		wait = func() { g.Wait() }
	default:
		return "The type of actual should be *sync.WaitGroup or coroutine.Group"
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
	case coroutine.Group:
		wait = func() { g.Wait() }
	default:
		return "The type of actual should be *sync.WaitGroup or coroutine.Group"
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
