package hashmap

var _ HashMap = (*normalHashMap)(nil)

type normalHashMap struct {
	ctx map[any]any
}

func (hm normalHashMap) Put(key HashKey, value any) HashMap {
	hm.ctx[key.Value()] = value

	return hm
}

func (hm normalHashMap) Get(key HashKey) (any, bool) {
	if value, ok := hm.ctx[key.Value()]; ok {
		return value, true
	}

	return nil, false
}

func (hm normalHashMap) GetString(key HashKey) (string, bool) {
	if value, ok := hm.ctx[key.Value()]; ok {
		return ToString(value), true
	}

	return emptyString, false
}

func (hm normalHashMap) Remove(key HashKey) HashMap {
	k := key.Value()
	if _, ok := hm.ctx[k]; ok {
		delete(hm.ctx, k)
	}

	return hm
}

func (hm normalHashMap) Has(key HashKey) bool {
	_, ok := hm.ctx[key.Value()]

	return ok
}

func (hm normalHashMap) Size() int {
	return len(hm.ctx)
}
