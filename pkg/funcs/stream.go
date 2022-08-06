package funcs

type (
	rxOptions struct {
		unlimitedWorkers bool
		workers          int
	}
	Option                     func(opts *rxOptions)
	KeyFunc[T any]             func(item T) T
	FilterFunc[T any]          func(item T) bool
	MapFunc[T any]             func(item T) T
	LessFunc[T any]            func(a, b T) bool
	WalkFunc[T any]            func(item T, pipe chan<- T)
	PredicateFunc[T any]       func(item T) bool
	ForAllFunc[T any]          func(pipe <-chan T)
	ForEachFunc[T any]         func(item T)
	ParallelFunc[T any]        func(item T)
	ReduceFunc[T any]          func(pipe <-chan T) (T, error)
	GenerateFunc[T any]        func(source <-chan T)
)

type Stream interface {
	Distinct(keyFunc KeyFunc) Stream
	Filter(filterFunc FilterFunc, opts ...Option) Stream
	Group(fn KeyFunc) Stream
	Head(n int64) Stream
	Tail(n int64) Stream
	First[T any]() T
	Last[T any]() T
	Map(fn MapFunc, opts ...Option) Stream
	Merge() Stream
	Reverse() Stream
	Sort(fn LessFunc) Stream
	Walk(fn WalkFunc, opts ...Option) Stream
	Concat(streams ...Stream) Stream
	AllMatch(fn PredicateFunc) bool
	AnyMatch(fn PredicateFunc) bool
	NoneMatch(fn PredicateFunc) bool
	ForAll(fn ForAllFunc)
	ForEach(fn ForEachFunc)
	Count() int64
	Done()
}
