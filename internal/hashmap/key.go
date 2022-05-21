package hashmap

import (
	"fmt"
)

type HashKey interface {
	Value() any
	HashCode() uint32
}

func ToString(src any) string {
	return fmt.Sprintf("%v", src)
}
