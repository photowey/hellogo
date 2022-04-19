package json

import (
	`strconv`
)

// ---------------------------------------------------------------- JSON

// JSOONObject {@code JSON} Object
type JSOONObject struct {
	context map[string]any
}

// ---------------------------------------------------------------- method

func (jsoon *JSOONObject) Put(key string, value any) {
	jsoon.context[key] = value
}

func (jsoon *JSOONObject) Get(key string) any {
	value, ok := jsoon.context[key]
	if ok {
		return value
	}

	return nil
}

func (jsoon *JSOONObject) GetSafe(key string, standBy any) (any, bool) {
	value, ok := jsoon.context[key]
	if ok {
		return value, true
	}

	return standBy, false
}

func (jsoon *JSOONObject) GetString(key string) (string, bool) {
	value, ok := jsoon.context[key]
	if ok {
		v, ook := value.(string)
		if ook {
			return v, true
		}
		return "", false
	}

	return "", false
}

func (jsoon *JSOONObject) GetInt64(key string) (int64, bool) {
	value, ok := jsoon.context[key]
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

func (jsoon *JSOONObject) GetFloat64(key string) (float64, bool) {
	value, ok := jsoon.context[key]
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

func (jsoon *JSOONObject) GetBool(key string) (bool, bool) {
	value, ok := jsoon.context[key]
	v, ok := value.(bool)
	if !ok {
		return false, false
	}

	return v, true
}

func (jsoon *JSOONObject) Remove(key string) {
	_, ok := jsoon.context[key]
	if ok {
		delete(jsoon.context, key)
	}
}

func (jsoon *JSOONObject) ToJSONString() (string, error) {
	return ToJSONString(jsoon.context)
}

// ---------------------------------------------------------------- function

func NewJsoonObject() *JSOONObject {
	return &JSOONObject{
		context: make(map[string]any),
	}
}

func NewJsoonObjects(mv map[string]any) *JSOONObject {
	mvp := NewJsoonObject()
	if len(mv) > 0 {
		for k, v := range mv {
			mvp.Put(k, v)
		}
	}

	return mvp
}

func ParseJSONObject(body string) (*JSOONObject, error) {
	mv := make(map[string]any)
	err := ToStruct([]byte(body), &mv)
	if err != nil {
		return nil, err
	}

	return NewJsoonObjects(mv), nil
}
