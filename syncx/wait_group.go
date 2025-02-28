package syncx

import (
	"context"
	"sync"
)

// WaitGroup helper for `sync.WaitGroup`
type WaitGroup struct {
	wg sync.WaitGroup
}

// Add delta, detail see `sync.WaitGroup.Add`
func (s *WaitGroup) Add(delta int) {
	s.wg.Add(delta)
}

// Done decrements the WaitGroup counter by one.
func (s *WaitGroup) Done() {
	s.wg.Done()
}

// Wait return a channel, the channel will close when the `sync.WaitGroup` counter is zero.
func (s *WaitGroup) Wait() <-chan struct{} {
	return s.watchDone()
}

// WaitOrDone will return when context done or wait group done
//
//   - nil: wait group done
//   - error: context done
func (s *WaitGroup) WaitOrDone(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-s.Wait():
		return nil
	}
}

func (s *WaitGroup) watchDone() chan struct{} {
	done := make(chan struct{})
	go func() {
		s.wg.Wait()
		close(done)
	}()

	return done
}
