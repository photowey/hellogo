package lists

import (
	"github.com/hellogo/pkg/generic/collections"
	"github.com/hellogo/pkg/generic/functions"
)

type List[T comparable] interface {
	collections.Collection[T]
	Sort(cpt functions.Comparator[T])
	Get(index int, t T) T
	Set(index int, t T)
	RemoveIdx(index int)
	IndexOf(t T) int
	LastIndexOf(t T) int
}
