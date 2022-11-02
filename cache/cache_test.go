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
// 3. 手动清理缓存
// 4. 设置过期时间
// 5. 修改过期时间

func TestCache(t *testing.T) {
	table := Cache("TestCache")
	// 测试写入
	table.Add(k+"_1", 0*time.Second, v)
	// 测试写入时间
	table.Add(k+"_2", 5*time.Second, v)

	p, err := table.Value(k + "_1")
	if err != nil || p == nil || p.Value.(string) != v {
		t.Error("读取没有设置过期时间的缓存失败:", err)
	}
	p, err = table.Value(k + "_2")
	if err != nil || p == nil || p.Value.(string) != v {
		t.Error("读取设置过期时间的缓存失败:", err)
	}
}

func TestCacheExpire(t *testing.T) {
	table := Cache("TestCacheExpire")

	table.Add(k+"_1", 3*time.Second, v)
	table.Add(k+"_2", 3*time.Second, v)

	time.Sleep(2 * time.Second)

	_, err := table.Value(k + "_1")
	if err != nil {
		t.Error("读取设置过期时间的缓存失败:", err)
	}
	time.Sleep(2 * time.Second)

	_, err = table.Value(k + "_1")
	if err != nil {
		t.Error("读取访问过的的缓存失败:", err)
	}

	_, err = table.Value(k + "_2")
	if err == nil {
		t.Error("初始过期时间失败:", err)
	}
}

func TestSetExpire(t *testing.T) {
	table := Cache("TestSetExpire")
	table.Add(k+"_e_1", 2000*time.Microsecond, v)
	time.Sleep(1500 * time.Microsecond)
	table.SetExpire(k+"_e_1", 2000*time.Microsecond)
	time.Sleep(1500 * time.Microsecond)
	_, err := table.Value(k + "_e_1")
	if err != nil {
		t.Error("重置过期时间失败:", err)
	}
}
