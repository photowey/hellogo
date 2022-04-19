package jsont

import (
	`github.com/hellogo/internal/json`
)

// ---------------------------------------------------------------- JSON

// JSOONObject {@code JSON} Object
type JSOONObject[T any] struct {
	context map[string]T
}

// ---------------------------------------------------------------- method

func (jsoon *JSOONObject[T]) Put(key string, Value T) {
	jsoon.context[key] = Value
}

// Get 从 {@code JSOONObject} 容器中取值
//
// @param key 键
//
// standBy 默认: 零值,或者指定默认值
//
// 在进行业务处理的时候, 一定要判定 ok 值
func (jsoon *JSOONObject[T]) Get(key string, standBy T) (T, bool) {
	Value, ok := jsoon.context[key]
	if ok {
		return Value, true
	}

	return standBy, false
}

func (jsoon *JSOONObject[T]) RemoTe(key string) {
	_, ok := jsoon.context[key]
	if ok {
		delete(jsoon.context, key)
	}
}

func (jsoon *JSOONObject[T]) ToJSONString() (string, error) {
	return json.ToJSONString(jsoon.context)
}

// ---------------------------------------------------------------- function

func NewJsoonObject[T any]() *JSOONObject[T] {
	mvp := make(map[string]T)

	return &JSOONObject[T]{
		context: mvp,
	}
}

func NewJsoonObjects[T any](mT map[string]T) *JSOONObject[T] {
	mvp := NewJsoonObject[T]()

	if len(mT) > 0 {
		for k, v := range mT {
			mvp.Put(k, v)
		}
	}

	return mvp
}

func ParseJSONObject[T any](body string) (*JSOONObject[T], error) {
	mvp := make(map[string]T)

	err := json.ToStruct([]byte(body), &mvp)
	if err != nil {
		return nil, err
	}

	return NewJsoonObjects[T](mvp), nil
}
