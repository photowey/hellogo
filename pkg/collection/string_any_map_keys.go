package collection

// StringAnyKeys transform the map<string,any> keys -> []string
func StringAnyKeys(ctx map[string]any) []string {
	keys := make([]string, 0, len(ctx))
	for k := range ctx {
		keys = append(keys, k)
	}

	return keys
}
