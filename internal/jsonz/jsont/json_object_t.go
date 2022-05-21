package jsont

import (
	"github.com/hellogo/internal/jsonz"
)

// ---------------------------------------------------------------- JSON

// Object {@code JSON} Object
type Object[T any] struct {
	ctx map[string]T
}

// ---------------------------------------------------------------- method

func (ob *Object[T]) Put(key string, Value T) {
	ob.ctx[key] = Value
}

// Get 从 {@code Object} 容器中取值
//
// @param key 键
//
// standBy 默认: 零值,或者指定默认值
//
// 在进行业务处理的时候, 一定要判定 ok 值
func (ob *Object[T]) Get(key string, standBy T) (T, bool) {
	Value, ok := ob.ctx[key]
	if ok {
		return Value, true
	}

	return standBy, false
}

func (ob *Object[T]) RemoTe(key string) {
	_, ok := ob.ctx[key]
	if ok {
		delete(ob.ctx, key)
	}
}

func (ob *Object[T]) String() (string, error) {
	return jsonz.String(ob.ctx)
}

// ---------------------------------------------------------------- function

func NewObject[T any]() *Object[T] {
	mvp := make(map[string]T)

	return &Object[T]{
		ctx: mvp,
	}
}

func NewObjectWithMap[T any](ctx map[string]T) *Object[T] {
	mvp := NewObject[T]()

	if len(ctx) > 0 {
		for k, v := range ctx {
			mvp.Put(k, v)
		}
	}

	return mvp
}

func ParseObject[T any](body string) (*Object[T], error) {
	mvp := make(map[string]T)

	err := jsonz.UnmarshalStruct([]byte(body), &mvp)
	if err != nil {
		return nil, err
	}

	return NewObjectWithMap[T](mvp), nil
}
