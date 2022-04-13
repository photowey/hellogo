package cmap

/**
 * [K,V]
 */

import (
	"sync"
)

type RWMap struct {
	m    map[any]any
	lock sync.RWMutex
}

func NewRWLockMap() *RWMap {
	m := make(map[any]any, 0)

	return &RWMap{m: m}
}

func (m *RWMap) Get(key any) (any, bool) {
	m.lock.RLock()
	v, ok := m.m[key]
	m.lock.RUnlock()

	return v, ok
}

func (m *RWMap) Put(key any, value any) {
	m.lock.Lock()
	m.m[key] = value
	m.lock.Unlock()
}

func (m *RWMap) Remove(key any) {
	m.lock.Lock()
	delete(m.m, key)
	m.lock.Unlock()
}
