package cmap

import (
	"sync"
)

/**
 * [K,V]
 */

type CMap struct {
	partitions []*kernelMap
	buckets    int
}

type kernelMap struct {
	m    map[any]any
	lock sync.RWMutex
}

type PartitionedKey interface {
	Value() any
	PartitionKey() int64
}

// -------------------------------------------------------------

func createKernelMap() *kernelMap {
	return &kernelMap{
		m: make(map[any]any),
	}
}

func (km *kernelMap) get(key PartitionedKey) (any, bool) {
	keyVal := key.Value()
	km.lock.RLock()
	v, ok := km.m[keyVal]
	km.lock.RUnlock()
	return v, ok
}

func (km *kernelMap) put(key PartitionedKey, v any) {
	keyVal := key.Value()
	km.lock.Lock()
	km.m[keyVal] = v
	km.lock.Unlock()
}

func (km *kernelMap) remove(key PartitionedKey) {
	keyVal := key.Value()
	km.lock.Lock()
	delete(km.m, keyVal)
	km.lock.Unlock()
}

// -------------------------------------------------------------

func (cm *CMap) calculatePartition(key PartitionedKey) *kernelMap {
	partitionID := key.PartitionKey() % int64(cm.buckets)
	return (*kernelMap)(cm.partitions[partitionID])
}

// -------------------------------------------------------------

func NewCMap(index int) *CMap {
	var partitions []*kernelMap
	for i := 0; i < index; i++ {
		partitions = append(partitions, createKernelMap())
	}
	return &CMap{partitions, index}
}

// -------------------------------------------------------------

func (cm *CMap) Get(key PartitionedKey) (any, bool) {
	return cm.calculatePartition(key).get(key)
}

func (cm *CMap) Put(key PartitionedKey, v any) {
	km := cm.calculatePartition(key)
	km.put(key, v)
}

func (cm *CMap) Remove(key PartitionedKey) {
	km := cm.calculatePartition(key)
	km.remove(key)
}
