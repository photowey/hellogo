package cmap

import (
	"github.com/hellogo/internal/hash"
)

type StringKey struct {
	value string
}

func (key *StringKey) Value() interface{} {
	return key.value
}

func (key *StringKey) PartitionKey() int64 {
	return int64(hash.MakeHash(key.value))
}

func WrapperStringKey(key string) *StringKey {
	return &StringKey{key}
}
