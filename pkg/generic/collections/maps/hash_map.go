package maps

type HashMap[K comparable, V comparable] struct {
	ctx map[K]V
}

func (hm HashMap[K, V]) Size() int {
	return len(hm.ctx)
}

func (hm HashMap[K, V]) IsEmpty() bool {
	return hm.Size() == 0
}

func (hm HashMap[K, V]) ContainsKey(k K) bool {
	_, ok := hm.ctx[k]

	return ok
}

func (hm HashMap[K, V]) ContainsValue(v V) bool {
	for _, hv := range hm.ctx {
		if hv == v {
			return true
		}
	}

	return false
}

func (hm HashMap[K, V]) Get(k K) (V, bool) {
	v, ok := hm.ctx[k]

	return v, ok
}

func (hm *HashMap[K, V]) Put(k K, v V) {
	hm.ctx[k] = v
}

func (hm *HashMap[K, V]) Remove(k K) V {
	v, ok := hm.Get(k)
	if ok {
		delete(hm.ctx, k)
	}

	return v
}

func (hm *HashMap[K, V]) PutAll(other Map[K, V]) {
	if other == nil {
		return
	}
	for _, k := range other.KeySet() {
		v, _ := other.Get(k)
		hm.ctx[k] = v
	}
}

func (hm *HashMap[K, V]) Clear() {
	m := make(map[K]V)
	hm.ctx = m
}

func (hm HashMap[K, V]) KeySet() []K {
	keys := make([]K, hm.Size())
	for k := range hm.ctx {
		keys = append(keys, k)
	}

	return keys
}

func (hm HashMap[K, V]) Values() []V {
	values := make([]V, hm.Size())
	for _, v := range hm.ctx {
		values = append(values, v)
	}

	return values
}

func NewHashMap[K comparable, V comparable]() Map[K, V] {
	m := make(map[K]V, 0)
	hm := &HashMap[K, V]{
		ctx: m,
	}

	return hm
}
