package jsonz

import (
	"strconv"
)

// ---------------------------------------------------------------- JSON

// Object {@code JSON} Object
type Object struct {
	ctx map[string]any
}

// ---------------------------------------------------------------- method

func (ob *Object) Put(key string, value any) {
	ob.ctx[key] = value
}

func (ob *Object) Get(key string) any {
	value, ok := ob.ctx[key]
	if ok {
		return value
	}

	return nil
}

func (ob *Object) GetSafe(key string, standBy any) (any, bool) {
	value, ok := ob.ctx[key]
	if ok {
		return value, true
	}

	return standBy, false
}

func (ob *Object) GetString(key string) (string, bool) {
	value, ok := ob.ctx[key]
	if ok {
		v, ook := value.(string)
		if ook {
			return v, true
		}
		return "", false
	}

	return "", false
}

func (ob *Object) GetInt64(key string) (int64, bool) {
	value, ok := ob.ctx[key]
	if ok {
		switch value.(type) {
		case int:
			return int64(value.(int)), true
		case int8:
			return int64(value.(int8)), true
		case int16:
			return int64(value.(int16)), true
		case int32:
			return int64(value.(int32)), true
		case int64:
			return value.(int64), true
		case uint8:
			return int64(value.(uint8)), true
		case uint16:
			return int64(value.(uint16)), true
		case uint32:
			return int64(value.(uint32)), true
		case uint64:
			return int64(value.(uint64)), true
		case float32:
			v := float64(value.(float32))
			return int64(v), true
		case float64:
			fv := value.(float64)
			return int64(fv), true
		}
	}

	return 0, false
}

func (ob *Object) GetFloat64(key string) (float64, bool) {
	value, ok := ob.ctx[key]
	fv := float64(0)
	if ok {
		switch value.(type) {
		case float64:
			fv = value.(float64)
		case float32:
			fv = float64(value.(float32))
		case int64:
			fv = float64(value.(int64))
		case int:
			fv = float64(value.(int))
		case string:
			fv, _ = strconv.ParseFloat(value.(string), 64)
		default:
			return fv, false
		}
	}

	return fv, true
}

func (ob *Object) GetBool(key string) (bool, bool) {
	value, ok := ob.ctx[key]
	v, ok := value.(bool)
	if !ok {
		return false, false
	}

	return v, true
}

func (ob *Object) Remove(key string) {
	_, ok := ob.ctx[key]
	if ok {
		delete(ob.ctx, key)
	}
}

func (ob *Object) String() (string, error) {
	return String(ob.ctx)
}

// ---------------------------------------------------------------- function

func NewObject() *Object {
	return &Object{
		ctx: make(map[string]any),
	}
}

func NewObjectWithMap(mv map[string]any) *Object {
	mvp := NewObject()
	if len(mv) > 0 {
		for k, v := range mv {
			mvp.Put(k, v)
		}
	}

	return mvp
}

func ParseObject(body string) (*Object, error) {
	mv := make(map[string]any)
	err := UnmarshalStruct([]byte(body), &mv)
	if err != nil {
		return nil, err
	}

	return NewObjectWithMap(mv), nil
}
