package predicate

func Test[T any](handler func(args T) bool, args T) bool {
	return handler(args)
}

func Tests[T any](handler func(args ...T) bool, args ...T) bool {
	return handler(args...)
}
