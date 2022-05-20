package castz

import (
	"time"

	"github.com/spf13/cast"
)

//
// 类型转换
//
// 方法名称-以 E 结尾的标识 有错误返回值
//

// ---------------------------------------------------------------- bool

// ToBoolB casts an interface to a bool type.
func ToBoolB(src any) (bool, bool) {
	v, err := cast.ToBoolE(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ---------------------------------------------------------------- int

// ToInt64B casts an interface to an int64 type.
func ToInt64B(src any) (int64, bool) {
	v, err := cast.ToInt64E(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ToInt32B casts an interface to an int32 type.
func ToInt32B(src any) (int32, bool) {
	v, err := cast.ToInt32E(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ToInt16B casts an interface to an int16 type.
func ToInt16B(src any) (int16, bool) {
	v, err := cast.ToInt16E(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ToInt8B casts an interface to an int8 type.
func ToInt8B(src any) (int8, bool) {
	v, err := cast.ToInt8E(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ToIntB casts an interface to an int type.
func ToIntB(src any) (int, bool) {
	v, err := cast.ToIntE(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ToUInt64B casts an interface to an uint64 type.
func ToUInt64B(src any) (uint64, bool) {
	v, err := cast.ToUint64E(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ToUInt32B casts an interface to an uint32 type.
func ToUInt32B(src any) (uint32, bool) {
	v, err := cast.ToUint32E(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ToUInt16B casts an interface to an uint16 type.
func ToUInt16B(src any) (uint16, bool) {
	v, err := cast.ToUint16E(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ToUInt8B casts an interface to an uint8 type.
func ToUInt8B(src any) (uint8, bool) {
	v, err := cast.ToUint8E(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ToUIntB casts an interface to an uint type.
func ToUIntB(src any) (uint, bool) {
	v, err := cast.ToUintE(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ---------------------------------------------------------------- float

// ToFloat64B casts an interface to a float64 type.
func ToFloat64B(src any) (float64, bool) {
	v, err := cast.ToFloat64E(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ToFloat32B casts an interface to a float32 type.
func ToFloat32B(src any) (float32, bool) {
	v, err := cast.ToFloat32E(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ---------------------------------------------------------------- string

// ToStringB casts an interface to a string type.
func ToStringB(src any) (string, bool) {
	v, err := cast.ToStringE(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ---------------------------------------------------------------- map<string,x>

// ToStringMapB casts an interface to a map[string]any type.
func ToStringMapB(src any) (map[string]any, bool) {
	v, err := cast.ToStringMapE(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ToStringMapStringB casts an interface to a map[string]string type.
func ToStringMapStringB(src any) (map[string]string, bool) {
	v, err := cast.ToStringMapStringE(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ToStringMapStringSliceB casts an interface to a map[string][]string type.
func ToStringMapStringSliceB(src any) (map[string][]string, bool) {
	v, err := cast.ToStringMapStringSliceE(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ToStringMapInt64B casts an interface to a map[string]int64 type.
func ToStringMapInt64B(src any) (map[string]int64, bool) {
	v, err := cast.ToStringMapInt64E(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ---------------------------------------------------------------- slice

// ToSliceB casts an interface to a []any type.
func ToSliceB(src any) ([]any, bool) {
	v, err := cast.ToSliceE(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ToStringSliceB casts an interface to a []string type.
func ToStringSliceB(src any) ([]string, bool) {
	v, err := cast.ToStringSliceE(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ToIntSliceB casts an interface to a []int type.
func ToIntSliceB(src any) ([]int, bool) {
	v, err := cast.ToIntSliceE(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ---------------------------------------------------------------- time

// ToTimeB casts an interface to a time.Time type.
func ToTimeB(src any) (time.Time, bool) {
	v, err := cast.ToTimeE(src)
	if err != nil {
		return v, false
	}

	return v, true
}
