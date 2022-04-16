package common

import (
	"sort"
)

// SortWrapper 对需要进行排序的目标进行包装
//
// 这样使得目标对象不需要自己手动实现 {@code Interface} 的三个接口
type SortWrapper[T comparable] struct {
	target  []T
	swapper func(T, T) bool
}

func (receiver SortWrapper[T]) Len() int {
	var array []T = receiver.target

	return len(array)
}

func (receiver SortWrapper[T]) Less(i, j int) bool {
	return receiver.swapper(receiver.target[i], receiver.target[j])
}

func (receiver SortWrapper[T]) Swap(i, j int) {
	receiver.target[i], receiver.target[j] = receiver.target[j], receiver.target[i]
}

// Sort 排序
//
// 采用 {@code SortWrapper} 包装
func Sort[T comparable](sorter []T, swapper func(T, T) bool) {
	sort.Sort(SortWrapper[T]{
		target:  sorter,
		swapper: swapper,
	})
}
