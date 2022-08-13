package sets

import (
	"github.com/hellogo/pkg/generic/collections"
)

type Set[T any] interface {
	collections.Collection[T]
}
