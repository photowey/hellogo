package maps

type Map[K comparable, V any] interface {
	Size() int
	IsEmpty() bool
	ContainsKey(k K) bool
	ContainsValue(v V) bool
	Get(k K) (V, bool)
	Put(k K, v V)
	Remove(k K) V
	PutAll(m Map[K, V])
	Clear()
	KeySet() []K
	Values() []V
	ForEachR(fx func(K, V))
	ForEachW(fx func(K, V))
}
