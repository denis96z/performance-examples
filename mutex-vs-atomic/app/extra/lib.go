package extra

import (
	"math/rand"
	"sync"
	"time"

	mva "performance-examples/mutex-vs-atomic"
)

const (
	NumRequestsTotal = 10000000
)

func MainFunc(n int, initMap func(), getItem func(uint32) mva.Item) time.Duration {
	initMap()

	wg := sync.WaitGroup{}
	wg.Add(n)

	t1 := time.Now()

	rand.Seed(100)
	for i := 0; i < n; i++ {
		go func() {
			for j := uint32(0); j < NumRequestsTotal/uint32(n); j++ {
				mva.ProcessItem(getItem(j % mva.NumItems))
			}

			wg.Done()
		}()
	}

	wg.Wait()

	t2 := time.Now()
	return t2.Sub(t1)
}
