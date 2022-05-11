package concurrenthashmap

type HashKey interface {
	Value() any
	HashCode() int64
}
