package concurrenthashmap

import (
	"fmt"
)

const (
	emptyString = ""
)

type ConcurrentHashMap interface {
	Put(key HashKey, value any) ConcurrentHashMap
	Get(key HashKey) (any, bool)
	GetString(key HashKey) (string, bool)
	Remove(key HashKey) ConcurrentHashMap
	Has(key HashKey) bool
	UnsafeSize() int64
	Size() int64
	Transfer() bool
}

func ToString(src any) string {
	return fmt.Sprintf("%v", src)
}
