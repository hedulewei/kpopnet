package kpopnet

import (
	"sync"
)

type Key int

const (
	profileCacheKey Key = iota
)

var (
	mu    sync.Mutex
	cache = make(map[Key]interface{}, 1)
)

func cached(key Key, makev func() (interface{}, error)) (v interface{}, err error) {
	mu.Lock()
	defer mu.Unlock()

	v, ok := cache[key]
	if ok {
		return
	}

	if v, err = makev(); err != nil {
		return
	}
	cache[key] = v
	return
}