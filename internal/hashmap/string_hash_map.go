package hashmap

// StringHashMap a map of k,v both string
type StringHashMap interface {
	Put(key string, value string) StringHashMap
	Get(key string) (string, bool)
	GetString(key string) (string, bool)
	Remove(key string) StringHashMap
	Has(key string) bool
	Size() int
}
