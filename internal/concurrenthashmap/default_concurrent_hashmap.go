package concurrenthashmap

import (
	"errors"
	"sync"
)

var _ ConcurrentHashMap = (*concurrentHashMap)(nil)

type partition []*kernelMap

type concurrentHashMap struct {
	partitions partition
	buckets    int
}

type kernelMap struct {
	ctx  map[any]any
	lock sync.RWMutex
}

// ----------------------------------------------------------------

func createKernelMap() *kernelMap {
	return &kernelMap{
		ctx: make(map[any]any),
	}
}

func (km *kernelMap) get(key HashKey) (any, bool) {
	keyVal := key.Value()
	km.lock.RLock()
	defer km.lock.RUnlock()
	v, ok := km.ctx[keyVal]

	return v, ok
}

func (km *kernelMap) put(key HashKey, v any) {
	keyVal := key.Value()
	km.lock.Lock()
	defer km.lock.Unlock()
	km.ctx[keyVal] = v
}

func (km *kernelMap) remove(key HashKey) {
	keyVal := key.Value()
	km.lock.Lock()
	defer km.lock.Unlock()
	delete(km.ctx, keyVal)
}

// ----------------------------------------------------------------

func (chm *concurrentHashMap) calculatePartition(key HashKey) *kernelMap {
	partitionId := key.HashCode() & (int64(chm.buckets) - 1)

	return chm.partitions[partitionId]
}

// ----------------------------------------------------------------

func NewConcurrentHashMap(capacity int) (*concurrentHashMap, error) {
	if capacity < 2 {
		return nil, errors.New("capacity must be an integer multiple of 2")
	}
	zero := capacity & (capacity - 1)
	if 0 != zero {
		return nil, errors.New("capacity must be an integer multiple of 2")
	}

	var partitions partition
	for i := 0; i < capacity; i++ {
		partitions = append(partitions, createKernelMap())
	}

	return &concurrentHashMap{partitions: partitions, buckets: capacity}, nil
}

// ----------------------------------------------------------------

func (chm *concurrentHashMap) Put(key HashKey, v any) ConcurrentHashMap {
	km := chm.calculatePartition(key)
	km.put(key, v)

	return chm
}

func (chm *concurrentHashMap) Get(key HashKey) (any, bool) {
	if v, ok := chm.calculatePartition(key).get(key); ok {
		return v, true
	}

	return nil, false
}

func (chm *concurrentHashMap) GetString(key HashKey) (string, bool) {
	if v, ok := chm.calculatePartition(key).get(key); ok {
		return ToString(v), true
	}

	return emptyString, false
}

func (chm *concurrentHashMap) Remove(key HashKey) ConcurrentHashMap {
	km := chm.calculatePartition(key)
	km.remove(key)

	return chm
}

func (chm *concurrentHashMap) Has(key HashKey) bool {
	km := chm.calculatePartition(key)
	_, ok := km.get(key)

	return ok
}

func (chm *concurrentHashMap) UnsafeSize() int64 {
	length := chm.length()

	return length
}

func (chm *concurrentHashMap) length() int64 {
	var length int64
	for _, partition := range chm.partitions {
		partition.lock.RLock()
		length += int64(len(partition.ctx))
		partition.lock.RUnlock()
	}

	return length
}

func (chm *concurrentHashMap) Size() int64 {
	panic("not support now")

	return 0
}

func (chm *concurrentHashMap) Transfer() bool {
	panic("not support now")

	return false
}
