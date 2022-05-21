package hashmap

import (
	"github.com/hellogo/internal/hash"
)

var _ HashKey = (*Uint32Key)(nil)

type Uint32Key struct {
	value    any    // value
	hashCode uint32 // hash code
}

func (uk Uint32Key) Value() any {
	return uk.value
}

func (uk Uint32Key) HashCode() uint32 {
	return uk.hashCode
}

func WrapperUint32Key(key any) Uint32Key {
	return Uint32Key{
		value:    key,
		hashCode: hash.MakeHash(ToString(key)),
	}
}
