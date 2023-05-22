package db_vs_inprocess

import (
	"fmt"
	"sync"

	"github.com/bradfitz/gomemcache/memcache"
)

type Item struct {
	Key   uint32
	Value string
}

const (
	NumItems = 20000
)

var (
	iMu  sync.Mutex
	iMap map[uint32]string

	mc *memcache.Client
)

//go:noinline
func InitProcess() {
	iMu.Lock()
	defer iMu.Unlock()

	iMap = make(map[uint32]string)
	for i := uint32(0); i < NumItems; i++ {
		iMap[i] = makeItemValue(i)
	}
}

//go:noinline
func InitMemcached() {
	mc = memcache.New("127.0.0.1:11211")
	for i := uint32(0); i < NumItems; i++ {
		mc.Set(&memcache.Item{
			Key:   makeMemcacheItemKey(i),
			Value: []byte(makeItemValue(i)),
		})
	}
}

func makeMemcacheItemKey(k uint32) string {
	return fmt.Sprintf("%d", k)
}

func makeItemValue(k uint32) string {
	return fmt.Sprintf("%016d", k)
}

//go:noinline
func GetDataFromProcess(k uint32) Item {
	iMu.Lock()
	defer iMu.Unlock()

	v, ok := iMap[k]
	if !ok {
		panic(!ok)
	}

	return Item{
		Key:   k,
		Value: v,
	}
}

//go:noinline
func GetDataFromMemcached(k uint32) Item {
	it, err := mc.Get(makeMemcacheItemKey(k))
	if err != nil {
		panic(err)
	}
	return Item{
		Key:   k,
		Value: string(it.Value),
	}
}

//go:noinline
func ProcessItem(Item) {}
