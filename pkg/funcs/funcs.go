package funcs

import (
	"github.com/hellogo/pkg/predicate"
)

func ForEach[T any](src []T, handler func(T)) {
	for _, item := range src {
		handler(item)
	}
}

func Map[D any, T any](src []D, mapper func(D) T) []T {
	mapped := make([]T, 0, len(src))
	for _, item := range src {
		mapped = append(mapped, mapper(item))
	}

	return mapped
}

func Filter[T any](src []T, handler func(T) bool) []T {
	filtered := make([]T, 0)
	for _, item := range src {
		if predicate.Test(handler, item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}

func Reduce[T any](src []T, reduce func(src T, item T) T, zero T) T {
	value := zero
	for _, item := range src {
		value = reduce(value, item)
	}

	return value
}

func Find[T any](src []T, handler func(T) bool) (T, bool) {
	for _, item := range src {
		if predicate.Test(handler, item) {
			return item, true
		}
	}

	var zero T
	return zero, false
}

func FindIndex[T any](src []T, handler func(T) bool) int {
	for idx, item := range src {
		if predicate.Test(handler, item) {
			return idx
		}
	}

	return -1
}

func IndexOf[T comparable](src []T, target T) int {
	for idx, item := range src {
		if item == target {
			return idx
		}
	}

	return -1
}

func LastIndexOf[T comparable](src []T, target T) int {
	for i := len(src) - 1; i >= 0; i-- {
		if src[i] == target {
			return i
		}
	}

	return -1
}
