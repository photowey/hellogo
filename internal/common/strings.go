package common

/**
 * strings
 */

// ---------------------------------------------------------------- Contain

func StringsContains(haystack []string, needle string) (index int) {
	index = -1
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			index = i
			return
		}
	}
	return
}
