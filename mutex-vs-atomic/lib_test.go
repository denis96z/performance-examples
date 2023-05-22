package mutex_vs_atomic

import (
	"math/rand"
	"testing"
)

func BenchmarkProcess(b *testing.B) {
	InitMutexMap()

	initBenchmark()
	for i := 0; i < b.N; i++ {
		doBenchmarkIteration(GetDataFromMutexMap)
	}
}

func BenchmarkMemcached(b *testing.B) {
	InitAtomicMap()

	initBenchmark()
	for i := 0; i < b.N; i++ {
		doBenchmarkIteration(GetDataFromAtomicMap)
	}
}

func initBenchmark() {
	rand.Seed(100)
}

func doBenchmarkIteration(getData func(uint32) Item) {
	k := rand.Uint32() % NumItems
	ProcessItem(getData(k))
}
