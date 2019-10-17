package sync

import (
	"gotest.tools/assert"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		for i := 0; i < 3; i++ {
			counter.Inc()
		}
		assertCount(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i:=0; i< wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		assertCount(t, counter, wantedCount)
	})
}

func assertCount(t *testing.T, counter *Counter, expected int) {
	t.Helper()
	msg := "\nexpected: %d \nactual:   %d"
	value := counter.Value()
	assert.Equal(
		t, expected, value,
		msg, expected, value)
}
