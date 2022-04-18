package option

import (
	"fmt"
	"reflect"
)

// ---------------------------------------------------------------- Optional 设计

// Optional 选项设计
type Optional[T any] struct {
	data      T
	present   bool
	valueType string
}

// ComparableOptional 可比较选项设计
//
// 相较于 {@code Optional[T]} 不需要手动传比较器
type ComparableOptional[T comparable] struct {
	Optional[T]
}

// OptionalEmpty 返回一个空的 {@code Optional[T any]}
func OptionalEmpty[T any](zero T) Optional[T] {
	// FIXME 思考:
	// FIXME 如果 不加 zero 形参， 直接赋值？
	// FIXME data:      zero,
	// FIXME data 直接赋值零值为: nil, 可能触发的 {@code BUG}

	var optional Optional[T]
	optional.data = zero
	optional.present = false
	optional.valueType = ""

	return optional
}

// OptionalOf 构造 {@code Optional}
func OptionalOf[T any](value T) Optional[T] {
	var optional Optional[T]
	optional.data = value
	optional.present = true
	optional.valueType = reflect.TypeOf(value).String()

	return optional
}

// String 将 {@code Optional[T]} 包装的目标对象解析为字符串
func (optional Optional[T]) String() string {
	if optional.present {
		str := fmt.Sprintf("%v", optional.data)
		return str
	} else {
		return ""
	}
}

// Get 获取 {@code Optional[T]} 里面包装的目标对象
func (optional Optional[T]) Get() T {
	return optional.data
}

// OrElse 获取 {@code Optional[T]} 里面包装的目标对象
//
// 如果:
//
// 1.被包装对象有值 -> 返回 -> optional.data
//
// 2.被包装对象没有值 -> 返回 -> standBy
func (optional Optional[T]) OrElse(standBy T) T {
	if optional.IsPresent() {
		return optional.data
	}

	return standBy
}

// IsPresent 判断 {@code Optional[T]} 是否有 "真" 值
func (optional Optional[T]) IsPresent() bool {
	return optional.present
}

// Equals 比较是否相当
//
// @param value 比较对象
//
// @param comparator 自定义比较器
func (optional Optional[T]) Equals(value T, compareTo func(T, T) bool) bool {
	if optional.present {
		return compareTo(value, optional.data)
	}

	return false
}

// Equals 可比较类型
func (optional ComparableOptional[T]) Equals(value T) bool {
	if optional.present {
		//return value == optional.data
		return reflect.DeepEqual(value, optional.data)
	}

	return false
}
