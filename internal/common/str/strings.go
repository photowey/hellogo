package str

import (
	"bytes"
	"strconv"
	"strings"
	"sync"
)

/**
 * strings
 */

// ---------------------------------------------------------------- Contains

var syncPool = sync.Pool{
	New: func() interface{} {
		buf := make([]byte, 0)
		return bytes.NewBuffer(buf)
	},
}

func ArrayContains(haystack []string, needle string) (index int) {
	index = -1
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			index = i
			return
		}
	}
	return
}

func JoinInt64(src []int64) string {
	if len(src) == 0 {
		return ""
	}
	if len(src) == 1 {
		return strconv.FormatInt(src[0], 10)
	}

	buf := syncPool.Get().(*bytes.Buffer)
	for _, num := range src {
		buf.WriteString(formatInt64(num))
		buf.WriteByte(',')
	}

	if buf.Len() > 0 {
		buf.Truncate(buf.Len() - 1)
	}
	group := buf.String()

	buf.Reset()
	syncPool.Put(buf)

	return group
}

func SplitInt(src string) ([]int64, error) {
	if src == "" {
		return nil, nil
	}
	chars := strings.Split(src, ",")
	nums := make([]int64, 0, len(chars))
	for _, char := range chars {
		num, err := parseInt64(char)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}

	return nums, nil
}

func formatInt64(num int64) string {
	return strconv.FormatInt(num, 10)
}

func parseInt64(src string) (int64, error) {
	return strconv.ParseInt(src, 10, 64)
}
