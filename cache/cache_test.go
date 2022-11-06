package cache

import (
	"testing"
	"time"
)

var (
	k = "testKey"
	v = "testValue"
)

// 要实现的功能
// 1. 写入缓存
// 2. 读取缓存

func TestCache(t *testing.T) {
	table := NewCacheTabel("TestCache")
	// 测试写入key
	keyOne := k + "_one"
	keyTwo := k + "_two"
	table.Add(keyOne, 0, v)
	// 测试写入时间
	table.Add(keyTwo, 5*time.Second, v)

	p, err := table.Value(keyOne)
	if err != nil || p == nil || p.Value.(string) != v {
		t.Error("读取没有设置过期时间的缓存失败:", err)
	}
	p, err = table.Value(keyTwo)
	if err != nil || p == nil || p.Value.(string) != v {
		t.Error("读取设置过期时间的缓存失败:", err)
	}
}

func TestCacheExpire(t *testing.T) {
	table := NewCacheTabel("TestCacheExpire")
	keyOne := k + "_1"
	keyTwo := k + "_2"
	table.Add(keyOne, 3*time.Second, v)
	table.Add(keyTwo, time.Second, v)

	time.Sleep(2 * time.Second)

	item, err := table.Value(keyOne)
	if err != nil {
		t.Error("读取设置过期时间的缓存失败:", err)
	}
	if item.Value.(string) != v {
		t.Errorf("读取失败， 期望:%v, 得到:%v", v, item.Value)
	}
	time.Sleep(2 * time.Second)

	item, err = table.Value(keyOne)
	if err != ErrorKeyFound {
		t.Error("读取访问过的的缓存失败:", err)
	}
	_, err = table.Value(keyTwo)
	if err != ErrorKeyFound {
		t.Error("过期生效失败:", err)
	}
	table.Add(keyTwo, time.Second, v)
	value, err := table.Value(keyTwo)
	if err != nil {
		t.Error("过期生效失败:", err)
	}
	if value.Value.(string) != v {
		t.Fatal()
	}

}
