package apc

import (
	"sync"
	"time"
)

var apcCache sync.Map

type DataObj struct {
	Value    any
	ExpireAt time.Time
}

// Get 根据key，获得对应的value
func Get(key string) (any, bool) {
	v, ok := apcCache.Load(key)
	if !ok {
		return nil, false
	}
	data := v.(*DataObj)
	if data.ExpireAt.IsZero() {
		return data.Value, true
	}
	if data.ExpireAt.Before(time.Now()) {
		apcCache.Delete(key)
		return nil, false
	}
	return data.Value, true
}

// Set 设置k-v expires 单位：s
func Set(key string, value any, expires int64) {
	data := &DataObj{
		Value: value,
	}
	if expires > 0 {
		duration := time.Duration(expires) * time.Second
		data.ExpireAt = time.Now().Add(duration)
	}
	apcCache.Store(key, data)
}

// Delete 删除k-v
func Delete(key string) {
	apcCache.Delete(key)
}

// Clear 清空本地缓存
func Clear() {
	apcCache.Range(func(key, value any) bool {
		apcCache.Delete(key)
		return true
	})
}
