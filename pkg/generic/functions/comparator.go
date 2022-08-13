package functions

type Comparator[T comparable] func(t T) int
