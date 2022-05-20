package collection

// NewStringSlice 初始化一个字符串类型的切片
func NewStringSlice(cap ...int) StringSlice {
	switch len(cap) {
	case 1:
		if cap[0] > 0 {
			return make(StringSlice, cap[0])
		}
	}

	return make(StringSlice, 0)
}

// ToStringSlice 将 string 可变参数列表转换为 StringSlice
func ToStringSlice(haystack ...string) StringSlice {
	array := make(StringSlice, len(haystack))
	for i, v := range haystack {
		array[i] = v
	}

	return array
}
