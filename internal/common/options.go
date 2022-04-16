package common

import (
	"fmt"
	"reflect"
)

// ---------------------------------------------------------------- Optional 设计

// Optional 选项设计
type Optional[T any] struct {
	Data      T
	Present   bool
	ValueType string
}

// ComparableOptional 可比较选项设计
//
// 相较于 {@code Optional[T]} 不需要手动串比较器
type ComparableOptional[T comparable] struct {
	Optional[T]
}

// OptionalEmpty 返回一个空的 {@code Optional[T any]}
func OptionalEmpty[T any](zero T) Optional[T] {
	// FIXME 思考:
	// FIXME 如果 不叫 zero 形参， 直接赋值？
	// FIXME Data:      zero,
	// FIXME Data 直接赋值零值为: nil, 可能触发的 {@code BUG}

	var optional Optional[T]
	optional.Data = zero
	optional.Present = false
	optional.ValueType = ""

	return optional
}

// OptionalOf 构造 {@code Optional}
func OptionalOf[T any](value T) Optional[T] {
	var optional Optional[T]
	optional.Data = value
	optional.Present = false
	optional.ValueType = reflect.TypeOf(value).String()

	return optional
}

// ToString 将 {@code Optional[T]} 包装的目标对象解析为字符串
func (optional Optional[T]) ToString() string {
	if optional.Present {
		str := fmt.Sprintf("%v", optional.Data)
		return str
	} else {
		return "empty"
	}
}

// Get 获取 {@code Optional[T]} 里面包装的目标对象
func (optional Optional[T]) Get() T {
	return optional.Data
}

// IsPresent 判断 {@code Optional[T]} 是否有 "真" 值
func (optional Optional[T]) IsPresent() bool {
	return optional.Present
}

// Equals 比较是否相当
//
// @param value 比较对象
//
// @param comparator 自定义比较器
func (optional Optional[T]) Equals(value T, compareTo func(T, T) bool) bool {
	if optional.Present {
		return compareTo(value, optional.Data)
	}

	return false
}

// Equals 可比较类型
func (optional ComparableOptional[T]) Equals(value T) bool {
	if optional.Present {
		// value == optional.Data
		return reflect.DeepEqual(value, optional.Data)
	}

	return false
}

// ---------------------------------------------------------------- Result 设计

// Result 结果
//
// 成功 Ok
//
// 失败 error
//
// e.g.:
//
/*
func Request[R any](r R) Result[R] {
	return Result[R]{
		Ok:     r,
		Failed: nil,
	}
}
*/
type Result[T any] struct {
	Ok     T
	Failed error
}
