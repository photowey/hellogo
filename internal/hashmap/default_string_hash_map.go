package hashmap

import (
	"errors"
)

var _ StringHashMap = (*defaultStringHashMap)(nil)

type defaultStringHashMap struct {
	ctx map[string]string
}

func (hm defaultStringHashMap) Put(key string, value string) StringHashMap {
	hm.ctx[key] = value

	return hm
}

func (hm defaultStringHashMap) Get(key string) (string, bool) {
	if value, ok := hm.ctx[key]; ok {
		return value, true
	}

	return emptyString, false
}

func (hm defaultStringHashMap) GetString(key string) (string, bool) {
	if value, ok := hm.ctx[key]; ok {
		return ToString(value), true
	}

	return emptyString, false
}

func (hm defaultStringHashMap) Remove(key string) StringHashMap {
	k := key
	if _, ok := hm.ctx[k]; ok {
		delete(hm.ctx, k)
	}

	return hm
}

func (hm defaultStringHashMap) Has(key string) bool {
	_, ok := hm.ctx[key]

	return ok
}

func (hm defaultStringHashMap) Size() int {
	return len(hm.ctx)
}

func NewStringHashMap(capacities ...int) (StringHashMap, error) {
	capacity := 16
	switch len(capacities) {
	case 1:
		capacity = capacities[0]
	}

	if capacity < 2 {
		return nil, errors.New("capacity must be an integer multiple of 2")
	}
	zero := capacity & (capacity - 1)
	if 0 != zero {
		return nil, errors.New("capacity must be an integer multiple of 2")
	}

	return &defaultStringHashMap{make(map[string]string, capacity)}, nil
}
