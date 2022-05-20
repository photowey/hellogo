package collection

// Int64Keys transform the map<int64,int64>  keys -> []int64
func Int64Keys(ctx map[int64]int64) []int64 {
	keys := make([]int64, 0, len(ctx))
	for k := range ctx {
		keys = append(keys, k)
	}

	return keys
}
