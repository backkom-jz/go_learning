package _map

import (
	"fmt"
	"sync"
	"testing"
)

func TestMapMutex(t *testing.T) {
	m := make(map[string]int)
	var mu sync.RWMutex

	// 写操作
	go func() {
		mu.Lock()
		m["key"] = 42
		t.Log(m)
		mu.Unlock()
	}()

	// 读操作
	go func() {
		mu.RLock()
		fmt.Println(m["key"])
		t.Log(m)
		mu.RUnlock()
	}()
}

func TestMapSyncMap(t *testing.T) {
	var sm sync.Map

	// 写操作
	sm.Store("key", 42)

	// 读操作
	if value, ok := sm.Load("key"); ok {
		fmt.Println(value)
	}

	// 删除操作
	sm.Delete("key")

	// 遍历
	sm.Store("key1", "value1")
	sm.Store("key2", "value2")
	sm.Range(func(key, value any) bool {
		fmt.Printf("%s: %v\n", key, value)
		return true // 返回 false 停止遍历
	})
}
