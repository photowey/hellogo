package hashmap

import (
	`errors`
)

const (
	emptyString = ""
)

var _ HashMap = (*uint32HashMap)(nil)

type uint32HashMap struct {
	ctx map[any]any
}

func (hm uint32HashMap) Put(key HashKey, value any) HashMap {
	hm.ctx[FormatKey(key)] = value

	return hm
}

func (hm uint32HashMap) Get(key HashKey) (any, bool) {
	if value, ok := hm.ctx[FormatKey(key)]; ok {
		return value, true
	}

	return nil, false
}

func (hm uint32HashMap) GetString(key HashKey) (string, bool) {
	if value, ok := hm.ctx[FormatKey(key)]; ok {
		return ToString(value), true
	}

	return emptyString, false
}

func (hm uint32HashMap) Remove(key HashKey) HashMap {
	k := FormatKey(key)
	if _, ok := hm.ctx[k]; ok {
		delete(hm.ctx, k)
	}

	return hm
}

func (hm uint32HashMap) Has(key HashKey) bool {
	_, ok := hm.ctx[FormatKey(key)]

	return ok
}

func (hm uint32HashMap) Size() int {
	return len(hm.ctx)
}

func NewHashMap(capacities ...int) (HashMap, error) {
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

	return &uint32HashMap{make(map[any]any, capacity)}, nil
}

func FormatKey(key HashKey) uint32 {
	return key.HashCode()
}
