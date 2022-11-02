package cache

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var defaultExpire = time.Minute

type CacheTable struct {
	sync.RWMutex
	// 缓存名称 每个table就是一个缓存系统
	name string
	// 缓存 对象
	items map[interface{}]*CacheItem
}

func (c *CacheTable) Add(key interface{}, expire time.Duration, value interface{}) *CacheItem {
	if expire == 0 {
		log.Printf("未设置时间")
		expire = defaultExpire
	}
	log.Printf("设置删除定时:%v,设置时间:%v", expire, time.Now())
	t := time.AfterFunc(expire, func() {
		c.DelItem(key)
	})
	item := NewCacheItem(key, expire, t, value)
	c.items[key] = item
	return item
}

func (c *CacheTable) Value(key interface{}) (*CacheItem, error) {
	if c.items == nil {
		return nil, ErrorKeyFound
	}
	if r, ok := c.items[key]; ok {
		s := r.cleanupTimer.Stop()
		if s {
			fmt.Println("停止删除")
		}
		log.Println("r.expireTime", r.expireTime)
		log.Printf("重新设置删除定时:%v, 设置时间:%v", r.expireTime, time.Now())
		r.cleanupTimer = time.AfterFunc(r.expireTime, func() { c.DelItem(key) })
		return r, nil
	}

	return nil, ErrorKeyFound
}

func (c *CacheTable) DelItem(key interface{}) error {
	log.Printf("执行删除定时:%v", time.Now())
	if c.items == nil {
		return nil
	}
	if _, ok := c.items[key]; ok {
		c.items[key] = nil
		return nil
	}
	return ErrorKeyFound
}

func (c *CacheTable) SetExpire(key interface{}, t time.Duration) error {
	if item, ok := c.items[key]; ok {
		//
		item.expireTime = t
		item.cleanupTimer.Reset(item.expireTime)
		return nil
	}
	return ErrorKeyFound
}
