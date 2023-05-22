package mutex_vs_atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Item struct {
	Key   uint32
	Value string
}

var (
	mu    sync.RWMutex
	muMap map[uint32]string

	aMap atomic.Pointer[map[uint32]string]
)

const (
	NumItems = 10000
)

func InitMutexMap() {
	muMap = makeMap()
}

func InitAtomicMap() {
	m := makeMap()
	aMap.Store(&m)
}

func makeMap() map[uint32]string {
	m := make(map[uint32]string)
	for i := uint32(0); i < NumItems; i++ {
		m[i] = makeItemValue(i)
	}
	return m
}

func makeItemValue(k uint32) string {
	return fmt.Sprintf("%016d", k)
}

//go:noinline
func GetDataFromMutexMap(k uint32) Item {
	mu.RLock()
	defer mu.RUnlock()

	return getItemFromMap(muMap, k)
}

//go:noinline
func GetDataFromAtomicMap(k uint32) Item {
	m := aMap.Load()
	return getItemFromMap(*m, k)
}

func getItemFromMap(m map[uint32]string, k uint32) Item {
	v, ok := m[k]
	if !ok {
		panic(!ok)
	}
	return Item{
		Key:   k,
		Value: v,
	}
}

//go:noinline
func ProcessItem(Item) {}
