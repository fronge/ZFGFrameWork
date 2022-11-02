package cache

import (
	"sync"
	"time"
)

type CacheItem struct {
	sync.RWMutex
	// 清空时间
	cleanupTimer *time.Timer
	// 过期时间
	expireTime time.Duration
	key        interface{}
	Value      interface{}
}

// NewCacheItem return a new CacheItem
// Parameter key is the item's cache-key.
// Parameter cleanupTimer is the item’s clearn Timer.
// Parameter expire is the item’s expire time.
// Parameter data is the item's cache-value.
func NewCacheItem(key interface{}, expire time.Duration, cleanupTimer *time.Timer, data interface{}) *CacheItem {
	return &CacheItem{
		cleanupTimer: cleanupTimer,
		expireTime:   expire,
		key:          key,
		Value:        data,
	}
}
