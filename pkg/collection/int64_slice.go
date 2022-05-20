package collection

// NewInt64Slice 初始化一个 int64 整型的切片
func NewInt64Slice(cap ...int) Int64Slice {
	switch len(cap) {
	case 1:
		if cap[0] > 0 {
			return make(Int64Slice, cap[0])
		}
	}
	return make(Int64Slice, 0)
}

// ToInt64Slice 将 int64 可变参数列表转换为 Int64Slice
func ToInt64Slice(haystack ...int64) Int64Slice {
	array := make(Int64Slice, len(haystack))
	for i, v := range haystack {
		array[i] = v
	}

	return array
}
