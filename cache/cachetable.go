package cache

import (
	"sync"
	"time"
)

var defaultExpire = time.Second
var cacheTables = map[string]*CacheTable{}

type CacheTable struct {
	sync.RWMutex
	// 缓存名称 每个table就是一个缓存系统
	name string
	// 缓存 对象
	items map[string]*CacheItem
}

func NewCacheTabel(name string) *CacheTable {
	if cachTable, ok := cacheTables[name]; ok && cachTable != nil {
		return cachTable
	}
	cacheTable := &CacheTable{
		name:  name,
		items: map[string]*CacheItem{},
	}
	cacheTables[name] = cacheTable

	return cacheTable
}

func (c *CacheTable) Add(key string, expire time.Duration, value interface{}) {
	if expire == 0 {
		expire = defaultExpire
	}
	// 判断是否存在
	if item, ok := c.items[key]; ok && item != nil {
		item.cleanupTimer.Stop()
	}
	// 定期删除
	t := time.AfterFunc(expire, func() {
		c.DelItem(key)
	})

	// item := NewCacheItem(key, expire, t, value)
	c.items[key] = NewCacheItem(key, expire, t, value)
	return
}

func (c *CacheTable) Value(key string) (*CacheItem, error) {
	if c.items == nil {
		return nil, ErrorKeyFound
	}
	if r, ok := c.items[key]; ok && r != nil {
		if time.Now().After(r.start.Add(r.expireTime)) {
			return nil, ErrorKeyFound
		}
		if s := r.cleanupTimer.Stop(); s {
			r.cleanupTimer = time.AfterFunc(r.expireTime, func() { c.DelItem(key) })
		}
		return r, nil
	}

	return nil, ErrorKeyFound
}

func (c *CacheTable) DelItem(key string) error {
	if c.items == nil {
		return nil
	}
	if _, ok := c.items[key]; ok {
		c.items[key] = nil
		return nil
	}
	return ErrorKeyFound
}
