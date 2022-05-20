package collection

// NewInt64Map new Int64Map
func NewInt64Map(cap ...int) Int64Map {
	switch len(cap) {
	case 1:
		if cap[0] > 0 {
			return make(Int64Map, cap[0])
		}
	}

	return make(Int64Map)
}

// InitInt64Map init Int64Map by given map
func InitInt64Map(ctx map[int64]int64) Int64Map {
	sm := make(Int64Map)
	if ctx != nil {
		for k, v := range ctx {
			sm[k] = v
		}
	}

	return sm
}

func (im Int64Map) Put(key, value int64) Int64Map {
	im[key] = value

	return im
}

func (im Int64Map) Get(key int64) (int64, bool) {
	value, ok := im[key]

	return value, ok
}

func (im Int64Map) Remove(key int64) {
	if im.Has(key) {
		delete(im, key)
	}
}

func (im Int64Map) Has(key int64) bool {
	_, ok := im[key]

	return ok
}

func (im Int64Map) Size() int {
	return len(im)
}
