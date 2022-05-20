package collection

import (
	"github.com/hellogo/pkg/castz"
)

// NewMixedMap new MixedMap
func NewMixedMap(cap ...int) MixedMap {
	switch len(cap) {
	case 1:
		if cap[0] > 0 {
			return make(MixedMap, cap[0])
		}
	}

	return make(MixedMap)
}

// InitMixedMap init MixedMap by given map
func InitMixedMap(ctx map[string]any) MixedMap {
	sm := make(MixedMap)
	if ctx != nil {
		for k, v := range ctx {
			sm[k] = v
		}
	}

	return sm
}

func (mm MixedMap) Put(key string, value any) MixedMap {
	mm[key] = value

	return mm
}

func (mm MixedMap) Get(key string) (any, bool) {
	value, ok := mm[key]

	return value, ok
}

func (mm MixedMap) Remove(key string) {
	if mm.Has(key) {
		delete(mm, key)
	}
}

func (mm MixedMap) Has(key string) bool {
	_, ok := mm[key]

	return ok
}

func (mm MixedMap) Size() int {
	return len(mm)
}

// ---------------------------------------------------------------- type value

func (mm MixedMap) GetString(key string) (string, bool) {
	if value, ok := mm[key]; ok {
		return castz.ToStringB(value)
	}

	return "", false
}

// ---------------------------------------------------------------- int

func (mm MixedMap) GetInt(key string) (int, bool) {
	if value, ok := mm[key]; ok {
		return castz.ToIntB(value)
	}

	return 0, false
}

func (mm MixedMap) GetInt64(key string) (int64, bool) {
	if value, ok := mm[key]; ok {
		return castz.ToInt64B(value)
	}

	return 0, false
}

func (mm MixedMap) GetInt32(key string) (int32, bool) {
	if value, ok := mm[key]; ok {
		return castz.ToInt32B(value)
	}

	return 0, false
}

func (mm MixedMap) GetInt16(key string) (int16, bool) {
	if value, ok := mm[key]; ok {
		return castz.ToInt16B(value)
	}

	return 0, false
}

func (mm MixedMap) GetInt8(key string) (int8, bool) {
	if value, ok := mm[key]; ok {
		return castz.ToInt8B(value)
	}

	return 0, false
}

// ---------------------------------------------------------------- uint

func (mm MixedMap) GetUInt(key string) (uint, bool) {
	if value, ok := mm[key]; ok {
		return castz.ToUIntB(value)
	}

	return 0, false
}

func (mm MixedMap) GetUInt64(key string) (uint64, bool) {
	if value, ok := mm[key]; ok {
		return castz.ToUInt64B(value)
	}

	return 0, false
}

func (mm MixedMap) GetUInt32(key string) (uint32, bool) {
	if value, ok := mm[key]; ok {
		return castz.ToUInt32B(value)
	}

	return 0, false
}

func (mm MixedMap) GetUInt16(key string) (uint16, bool) {
	if value, ok := mm[key]; ok {
		return castz.ToUInt16B(value)
	}

	return 0, false
}

func (mm MixedMap) GetUInt8(key string) (uint8, bool) {
	if value, ok := mm[key]; ok {
		return castz.ToUInt8B(value)
	}

	return 0, false
}

// ---------------------------------------------------------------- float

func (mm MixedMap) GetFloat64(key string) (float64, bool) {
	if value, ok := mm[key]; ok {
		return castz.ToFloat64B(value)
	}

	return 0, true
}

func (mm MixedMap) GetFloat32(key string) (float32, bool) {
	if value, ok := mm[key]; ok {
		return castz.ToFloat32B(value)
	}

	return 0, true
}

// ---------------------------------------------------------------- bool

func (mm MixedMap) GetBool(key string) (bool, bool) {
	if value, ok := mm[key]; ok {
		return castz.ToBoolB(value)
	}

	return false, false
}
