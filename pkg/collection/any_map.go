package collection

import (
	"github.com/hellogo/pkg/castz"
)

// NewAnyMap new AnyMap
func NewAnyMap() AnyMap {
	return make(AnyMap)
}

// InitAnyMap init AnyMap by given map
func InitAnyMap(ctx map[any]any) AnyMap {
	am := make(AnyMap)
	if ctx != nil {
		for k, v := range ctx {
			am[k] = v
		}
	}

	return am
}

func (am AnyMap) Put(key, value any) AnyMap {
	am[key] = value

	return am
}

func (am AnyMap) Get(key any) (any, bool) {
	value, ok := am[key]

	return value, ok
}

func (am AnyMap) Remove(key any) {
	if am.Has(key) {
		delete(am, key)
	}
}

func (am AnyMap) Has(key any) bool {
	_, ok := am[key]

	return ok
}

func (am AnyMap) Size() int {
	return len(am)
}

// ---------------------------------------------------------------- type value

func (am AnyMap) GetString(key string) (string, bool) {
	if value, ok := am[key]; ok {
		return castz.ToStringB(value)
	}

	return "", false
}

// ---------------------------------------------------------------- int

func (am AnyMap) GetInt(key string) (int, bool) {
	if value, ok := am[key]; ok {
		return castz.ToIntB(value)
	}

	return 0, false
}

func (am AnyMap) GetInt64(key string) (int64, bool) {
	if value, ok := am[key]; ok {
		return castz.ToInt64B(value)
	}

	return 0, false
}

func (am AnyMap) GetInt32(key string) (int32, bool) {
	if value, ok := am[key]; ok {
		return castz.ToInt32B(value)
	}

	return 0, false
}

func (am AnyMap) GetInt16(key string) (int16, bool) {
	if value, ok := am[key]; ok {
		return castz.ToInt16B(value)
	}

	return 0, false
}

func (am AnyMap) GetInt8(key string) (int8, bool) {
	if value, ok := am[key]; ok {
		return castz.ToInt8B(value)
	}

	return 0, false
}

// ---------------------------------------------------------------- uint

func (am AnyMap) GetUInt(key string) (uint, bool) {
	if value, ok := am[key]; ok {
		return castz.ToUIntB(value)
	}

	return 0, false
}

func (am AnyMap) GetUInt64(key string) (uint64, bool) {
	if value, ok := am[key]; ok {
		return castz.ToUInt64B(value)
	}

	return 0, false
}

func (am AnyMap) GetUInt32(key string) (uint32, bool) {
	if value, ok := am[key]; ok {
		return castz.ToUInt32B(value)
	}

	return 0, false
}

func (am AnyMap) GetUInt16(key string) (uint16, bool) {
	if value, ok := am[key]; ok {
		return castz.ToUInt16B(value)
	}

	return 0, false
}

func (am AnyMap) GetUInt8(key string) (uint8, bool) {
	if value, ok := am[key]; ok {
		return castz.ToUInt8B(value)
	}

	return 0, false
}

// ---------------------------------------------------------------- float

func (am AnyMap) GetFloat64(key string) (float64, bool) {
	if value, ok := am[key]; ok {
		return castz.ToFloat64B(value)
	}

	return 0, true
}

func (am AnyMap) GetFloat32(key string) (float32, bool) {
	if value, ok := am[key]; ok {
		return castz.ToFloat32B(value)
	}

	return 0, true
}

// ---------------------------------------------------------------- bool

func (am AnyMap) GetBool(key string) (bool, bool) {
	if value, ok := am[key]; ok {
		return castz.ToBoolB(value)
	}

	return false, false
}
