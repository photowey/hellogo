package collection

// Int64AnyKeys transform the map<int64,any> keys -> []int64
func Int64AnyKeys(ctx map[int64]any) []int64 {
	keys := make([]int64, 0, len(ctx))
	for k := range ctx {
		keys = append(keys, k)
	}

	return keys
}
