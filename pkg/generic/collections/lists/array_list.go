package lists

import (
	"github.com/hellogo/pkg/generic/collections"
	"github.com/hellogo/pkg/generic/collections/iterators"
	"github.com/hellogo/pkg/generic/functions"
)

var _ List[any] = (*ArrayList[any])(nil)

type ArrayList[T comparable] struct{}

func (al *ArrayList[T]) Iterator() iterators.Iterator[T] {
	return nil
}

func (al *ArrayList[T]) ForEach(action functions.Consumer[T]) {
}

func (al *ArrayList[T]) Size() int {
	return 0
}

func (al *ArrayList[T]) IsEmpty() bool {
	return false
}

func (al *ArrayList[T]) Contains(e T) bool {
	return false
}

func (al *ArrayList[T]) ContainsAll(c collections.Collection[T]) bool {
	return false
}

func (al *ArrayList[T]) ToArray() []T {
	var array []T

	return array
}

func (al *ArrayList[T]) Add(e T) bool {
	return false
}

func (al *ArrayList[T]) AddAll(c collections.Collection[T]) bool {
	return false
}

func (al *ArrayList[T]) Remove(e T) bool {
	return false
}

func (al *ArrayList[T]) RemoveAll(c collections.Collection[T]) bool {
	return false
}

func (al *ArrayList[T]) RemoveIf(filter functions.Predicate[T]) bool {
	return false
}

func (al *ArrayList[T]) Clear() {
}

func (al *ArrayList[T]) Sort(cpt functions.Comparator[T]) {
}

func (al *ArrayList[T]) Get(index int, t T) T {
	var dt T
	dt = t

	return dt
}

func (al *ArrayList[T]) Set(index int, t T) {
}

func (al *ArrayList[T]) RemoveIdx(index int) {
}

func (al *ArrayList[T]) IndexOf(t T) int {
	return 0
}

func (al *ArrayList[T]) LastIndexOf(t T) int {
	return 0
}
