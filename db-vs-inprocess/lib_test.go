package db_vs_inprocess

import (
	"math/rand"
	"testing"
)

func BenchmarkProcess(b *testing.B) {
	InitProcess()

	initBenchmark()
	for i := 0; i < b.N; i++ {
		doBenchmarkIteration(GetDataFromProcess)
	}
}

func BenchmarkMemcached(b *testing.B) {
	InitMemcached()

	initBenchmark()
	for i := 0; i < b.N; i++ {
		doBenchmarkIteration(GetDataFromMemcached)
	}
}

func initBenchmark() {
	rand.Seed(100)
}

func doBenchmarkIteration(getData func(uint32) Item) {
	k := rand.Uint32() % NumItems
	ProcessItem(getData(k))
}
