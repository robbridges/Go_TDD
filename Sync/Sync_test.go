package Sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Increment counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCounter(t, counter, 3)
	})
	t.Run("It needs to run safe concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})

}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.GetValue() != want {
		t.Errorf("The counter should have been %d, but it was %d", want, got.GetValue())
	}
}

func NewCounter() *Counter {
	return &Counter{}
}
