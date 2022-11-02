package cache

import "sync"

var (
	cache = make(map[string]*CacheTable)
	mutex sync.RWMutex
)

func Cache(table string) *CacheTable {
	mutex.RLock()
	defer mutex.RUnlock()
	if t, ok := cache[table]; ok {
		return t
	}
	t := &CacheTable{
		name:  table,
		items: make(map[interface{}]*CacheItem),
	}
	return t

}
