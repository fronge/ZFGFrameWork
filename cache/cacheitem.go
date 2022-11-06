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
	// 创建时间
	start time.Time
	key   interface{}
	Value interface{}
}

func NewCacheItem(key interface{}, expire time.Duration, cleanupTimer *time.Timer, data interface{}) *CacheItem {
	return &CacheItem{
		cleanupTimer: cleanupTimer,
		expireTime:   expire,
		key:          key,
		start:        time.Now(),
		Value:        data,
	}
}
