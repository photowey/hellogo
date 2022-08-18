package maps

import (
	"reflect"
	"sync"
)

type HashMap[K comparable, V any] struct {
	sync.RWMutex
	ctx map[K]V
}

func (hm *HashMap[K, V]) Size() int {
	return len(hm.ctx)
}

func (hm *HashMap[K, V]) IsEmpty() bool {
	return hm.Size() == 0
}

func (hm *HashMap[K, V]) ContainsKey(k K) bool {
	hm.RLock()
	defer hm.RUnlock()
	_, ok := hm.ctx[k]

	return ok
}

func (hm *HashMap[K, V]) ContainsValue(v V) bool {
	hm.RLock()
	defer hm.RUnlock()

	for _, hv := range hm.ctx {
		if reflect.DeepEqual(hv, v) {
			return true
		}
	}

	return false
}

func (hm *HashMap[K, V]) Get(k K) (V, bool) {
	hm.RLock()
	defer hm.RUnlock()

	v, ok := hm.ctx[k]

	return v, ok
}

func (hm *HashMap[K, V]) Put(k K, v V) {
	hm.Lock()
	defer hm.Unlock()

	hm.ctx[k] = v
}

func (hm *HashMap[K, V]) Remove(k K) V {
	hm.RLock()
	v, ok := hm.Get(k)
	hm.RUnlock()

	if ok {
		hm.Lock()
		delete(hm.ctx, k)
		hm.Unlock()
	}

	return v
}

func (hm *HashMap[K, V]) PutAll(other Map[K, V]) {
	hm.Lock()
	defer hm.Unlock()

	if other == nil {
		return
	}
	for _, k := range other.KeySet() {
		v, _ := other.Get(k)
		hm.ctx[k] = v
	}
}

func (hm *HashMap[K, V]) Clear() {
	hm.Lock()
	defer hm.Unlock()
	
	m := make(map[K]V)
	hm.ctx = m
}

func (hm *HashMap[K, V]) KeySet() []K {
	hm.RLock()
	defer hm.RUnlock()

	keys := make([]K, 0, len(hm.ctx))
	for k := range hm.ctx {
		keys = append(keys, k)
	}

	return keys
}

func (hm *HashMap[K, V]) Values() []V {
	hm.RLock()
	defer hm.RUnlock()

	values := make([]V, 0, len(hm.ctx))
	for _, v := range hm.ctx {
		values = append(values, v)
	}

	return values
}

func (hm *HashMap[K, V]) ForEachR(fx func(K, V)) {
	hm.RLock()
	defer hm.RUnlock()

	for k, v := range hm.ctx {
		fx(k, v)
	}
}

func (hm *HashMap[K, V]) ForEachW(fx func(K, V)) {
	hm.Lock()
	defer hm.Unlock()

	for k, v := range hm.ctx {
		fx(k, v)
	}
}

func NewHashMap[K comparable, V any]() Map[K, V] {
	m := make(map[K]V, 0)
	hm := &HashMap[K, V]{
		ctx: m,
	}

	return hm
}
