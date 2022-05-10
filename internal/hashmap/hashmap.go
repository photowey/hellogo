package hashmap

type HashMap interface {
	Put(key HashKey, value any) HashMap
	Get(key HashKey) (any, bool)
	GetString(key HashKey) (string, bool)
	Remove(key HashKey) HashMap
	Has(key HashKey) bool
	Size() int
}
