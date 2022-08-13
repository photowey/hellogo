package collections

import (
	"github.com/hellogo/pkg/generic/collections/iterators"
	"github.com/hellogo/pkg/generic/functions"
)

type Collection[E any] interface {
	iterators.Iterable[E]
	Size() int
	IsEmpty() bool
	Contains(e E) bool
	ContainsAll(c Collection[E]) bool
	ToArray() []E
	Add(e E) bool
	AddAll(c Collection[E]) bool
	Remove(e E) bool
	RemoveAll(c Collection[E]) bool
	RemoveIf(filter functions.Predicate[E]) bool
	Clear()
}
