package collection

//
// 集合类型
//

// ---------------------------------------------------------------- map

type (
	AnyMap    map[any]any
	Int64Map  map[int64]int64
	MixedMap  map[string]any
	StringMap map[string]string
)

// ---------------------------------------------------------------- string

type (
	StringSlice []string
	Int64Slice  []int64
)
