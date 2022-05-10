package hashmap

var _ HashKey = (*NormalKey)(nil)

type NormalKey struct {
	value    any    // value
	hashCode uint32 // hash code
}

func (uk NormalKey) Value() any {
	return uk.value
}

func (uk NormalKey) HashCode() uint32 {
	return uk.hashCode
}

func WrapperNormalKey(key any) Uint32Key {
	return Uint32Key{
		value:    key,
		hashCode: 0, // default value
	}
}
