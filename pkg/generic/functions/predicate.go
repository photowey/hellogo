package functions

type Predicate[T any] func(t T) bool

type PredicateAnd[T any] func(other T) Predicate[T]

type PredicateOr[T any] func(other T) Predicate[T]

type PredicateNegate[T any] func() Predicate[T]
