package concurrenthashmap

type Int64Key struct {
	value int64
}

func (key *Int64Key) Value() any {
	return key.value
}

func (key *Int64Key) HashCode() int64 {
	return key.value // value as hash code
}

func WrapperI64Key(key int64) *Int64Key {
	return &Int64Key{key}
}
