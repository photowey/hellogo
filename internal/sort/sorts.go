package sort

import (
	"sort"
)

// Wrapper 对需要进行排序的目标进行包装
//
// 这样使得目标对象不需要自己手动实现 {@code Interface} 的三个接口
type Wrapper[T comparable] struct {
	target  []T
	swapper func(T, T) bool
}

func (receiver Wrapper[T]) Len() int {
	var array []T = receiver.target

	return len(array)
}

func (receiver Wrapper[T]) Less(i, j int) bool {
	return receiver.swapper(receiver.target[i], receiver.target[j])
}

func (receiver Wrapper[T]) Swap(i, j int) {
	receiver.target[i], receiver.target[j] = receiver.target[j], receiver.target[i]
}

// Sort 排序
//
// 采用 {@code Wrapper} 包装
func Sort[T comparable](sorter []T, swapper func(T, T) bool) {
	sort.Sort(Wrapper[T]{
		target:  sorter,
		swapper: swapper,
	})
}
