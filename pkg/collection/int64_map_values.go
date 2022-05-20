package collection

// Int64Values transform the map<int64,int64> values []int64
func Int64Values(ctx map[int64]int64) []int64 {
	values := make([]int64, 0, len(ctx))
	for _, v := range ctx {
		values = append(values, v)
	}

	return values
}
